// Copyright 2014 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package vm

import (
	"github.com/holiman/uint256"

	"github.com/ledgerwatch/turbo-geth/common"
)

// fill in segment of operation information array for a block
func analyzeBlock(ctx *callCtx, pc uint64) (*BlockInfo, error) {
	blockInfo := NewBlockInfo(pc)
	ctx.contract.opsInfo[pc] = blockInfo
	code := ctx.contract.Code
	codeLen := len(code)
	jumpTable := ctx.interpreter.jt

	height := 0
	minHeight := 0
	maxHeight := 0
	for ; pc < uint64(codeLen); pc++ {
		op := OpCode(code[pc])
		oper := jumpTable[op]
		if oper == nil {
			continue
		}

		// track low and high watermark relative to block entry
		height -= oper.numPop
		minHeight = min(minHeight, height)	// will be <= 0
		height += oper.numPush
		maxHeight = max(maxHeight, height)	// will be >= 0
		blockInfo.constantGas += oper.constantGas

		if PUSH1 <= op && op <= PUSH32 {
		    pushByteSize := int(op) - int(PUSH1) + 1

			startMin := int(pc + 1)
			if startMin >= codeLen {
				startMin = codeLen
			}
			endMin := startMin + pushByteSize
			if startMin+pushByteSize >= codeLen {
				endMin = codeLen
			}

			integer := new(uint256.Int)
			integer.SetBytes(common.RightPadBytes(
				// So it doesn't matter what we push onto the stack.
				code[startMin:endMin], pushByteSize))

			// attach PushInfo with decoded push data to PUSHn
			pushInfo := NewPushInfo(pc, *integer)
			ctx.contract.opsInfo[pc] = pushInfo

			continue
		}
		
		// check jump destinations and optimize static jumps
		if op == JUMP || op == JUMPI {
			prevPC := pc - 1
			prevOp := OpCode(code[prevPC])
			if prevOp >= PUSH1 && prevOp <= PUSH32 {

				pos := ctx.contract.opsInfo[prevPC].(PushInfo).data
				if valid, _ := ctx.contract.validJumpdest(&pos); !valid {
					return nil, ErrInvalidJump
				}
				jumpInfo := NewJumpInfo(pc, pos.Uint64())
				ctx.contract.opsInfo[pc] = jumpInfo
				ctx.contract.opsInfo[prevPC] = nil

				// replace with JMP NOOP or JMPI NOOP and attach JumpInfo to JMP or JMPI
				if op == JUMP {
					code[prevPC] = byte(JMP)
				}
				if op == JUMPI {
					code[prevPC] = byte(JMPI)
				}
				code[pc] = byte(NOOP)

				// end block
				break
			}
		}
		if	op == JUMPDEST || op == STOP || op == RETURN || op == REVERT || op == SELFDESTRUCT {

			// end block
			break
		}
	}

	// min and max absolute stack length to avoid stack underflow or underflow
	blockInfo.minStack = -minHeight
	blockInfo.maxStack = maxHeight

	return blockInfo, nil
}

// codeBitmap collects data locations in code.
func codeBitmap(code []byte) []uint64 {
	// The bitmap is 4 bytes longer than necessary, in case the code
	// ends with a PUSH32, the algorithm will push zeroes onto the
	// bitvector outside the bounds of the actual code.
	bits := make([]uint64, (len(code)+32+63)/64)

	for pc := 0; pc < len(code); {
		op := OpCode(code[pc])
		pc++
		if op >= PUSH1 && op <= PUSH32 {
			numbits := int(op - PUSH1 + 1)
			x := uint64(1) << (op - PUSH1)
			x = x | (x - 1) // Smear the bit to the right
			idx := pc / 64
			shift := pc & 63
			bits[idx] |= x << shift
			if shift+shift > 64 {
				bits[idx+1] |= x >> (64 - shift)
			}
			pc += numbits
		}
	}
	return bits
}
