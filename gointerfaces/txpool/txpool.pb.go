// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: txpool/txpool.proto

package txpool

import (
	types "github.com/ledgerwatch/turbo-geth/gointerfaces/types"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ImportResult int32

const (
	ImportResult_SUCCESS        ImportResult = 0
	ImportResult_ALREADY_EXISTS ImportResult = 1
	ImportResult_FEE_TOO_LOW    ImportResult = 2
	ImportResult_STALE          ImportResult = 3
	ImportResult_INVALID        ImportResult = 4
	ImportResult_INTERNAL_ERROR ImportResult = 5
)

// Enum value maps for ImportResult.
var (
	ImportResult_name = map[int32]string{
		0: "SUCCESS",
		1: "ALREADY_EXISTS",
		2: "FEE_TOO_LOW",
		3: "STALE",
		4: "INVALID",
		5: "INTERNAL_ERROR",
	}
	ImportResult_value = map[string]int32{
		"SUCCESS":        0,
		"ALREADY_EXISTS": 1,
		"FEE_TOO_LOW":    2,
		"STALE":          3,
		"INVALID":        4,
		"INTERNAL_ERROR": 5,
	}
)

func (x ImportResult) Enum() *ImportResult {
	p := new(ImportResult)
	*p = x
	return p
}

func (x ImportResult) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ImportResult) Descriptor() protoreflect.EnumDescriptor {
	return file_txpool_txpool_proto_enumTypes[0].Descriptor()
}

func (ImportResult) Type() protoreflect.EnumType {
	return &file_txpool_txpool_proto_enumTypes[0]
}

func (x ImportResult) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ImportResult.Descriptor instead.
func (ImportResult) EnumDescriptor() ([]byte, []int) {
	return file_txpool_txpool_proto_rawDescGZIP(), []int{0}
}

type TxHashes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hashes []*types.H256 `protobuf:"bytes,1,rep,name=hashes,proto3" json:"hashes,omitempty"`
}

func (x *TxHashes) Reset() {
	*x = TxHashes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_txpool_txpool_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxHashes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxHashes) ProtoMessage() {}

func (x *TxHashes) ProtoReflect() protoreflect.Message {
	mi := &file_txpool_txpool_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxHashes.ProtoReflect.Descriptor instead.
func (*TxHashes) Descriptor() ([]byte, []int) {
	return file_txpool_txpool_proto_rawDescGZIP(), []int{0}
}

func (x *TxHashes) GetHashes() []*types.H256 {
	if x != nil {
		return x.Hashes
	}
	return nil
}

type ImportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Txs [][]byte `protobuf:"bytes,1,rep,name=txs,proto3" json:"txs,omitempty"`
}

func (x *ImportRequest) Reset() {
	*x = ImportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_txpool_txpool_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportRequest) ProtoMessage() {}

func (x *ImportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_txpool_txpool_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportRequest.ProtoReflect.Descriptor instead.
func (*ImportRequest) Descriptor() ([]byte, []int) {
	return file_txpool_txpool_proto_rawDescGZIP(), []int{1}
}

func (x *ImportRequest) GetTxs() [][]byte {
	if x != nil {
		return x.Txs
	}
	return nil
}

type ImportReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Imported []ImportResult `protobuf:"varint,1,rep,packed,name=imported,proto3,enum=txpool.ImportResult" json:"imported,omitempty"`
}

func (x *ImportReply) Reset() {
	*x = ImportReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_txpool_txpool_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportReply) ProtoMessage() {}

func (x *ImportReply) ProtoReflect() protoreflect.Message {
	mi := &file_txpool_txpool_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportReply.ProtoReflect.Descriptor instead.
func (*ImportReply) Descriptor() ([]byte, []int) {
	return file_txpool_txpool_proto_rawDescGZIP(), []int{2}
}

func (x *ImportReply) GetImported() []ImportResult {
	if x != nil {
		return x.Imported
	}
	return nil
}

type GetTransactionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hashes []*types.H256 `protobuf:"bytes,1,rep,name=hashes,proto3" json:"hashes,omitempty"`
}

func (x *GetTransactionsRequest) Reset() {
	*x = GetTransactionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_txpool_txpool_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTransactionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransactionsRequest) ProtoMessage() {}

func (x *GetTransactionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_txpool_txpool_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransactionsRequest.ProtoReflect.Descriptor instead.
func (*GetTransactionsRequest) Descriptor() ([]byte, []int) {
	return file_txpool_txpool_proto_rawDescGZIP(), []int{3}
}

func (x *GetTransactionsRequest) GetHashes() []*types.H256 {
	if x != nil {
		return x.Hashes
	}
	return nil
}

type GetTransactionsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Txs [][]byte `protobuf:"bytes,1,rep,name=txs,proto3" json:"txs,omitempty"`
}

func (x *GetTransactionsReply) Reset() {
	*x = GetTransactionsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_txpool_txpool_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTransactionsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransactionsReply) ProtoMessage() {}

func (x *GetTransactionsReply) ProtoReflect() protoreflect.Message {
	mi := &file_txpool_txpool_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransactionsReply.ProtoReflect.Descriptor instead.
func (*GetTransactionsReply) Descriptor() ([]byte, []int) {
	return file_txpool_txpool_proto_rawDescGZIP(), []int{4}
}

func (x *GetTransactionsReply) GetTxs() [][]byte {
	if x != nil {
		return x.Txs
	}
	return nil
}

type PendingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PendingRequest) Reset() {
	*x = PendingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_txpool_txpool_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PendingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PendingRequest) ProtoMessage() {}

func (x *PendingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_txpool_txpool_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PendingRequest.ProtoReflect.Descriptor instead.
func (*PendingRequest) Descriptor() ([]byte, []int) {
	return file_txpool_txpool_proto_rawDescGZIP(), []int{5}
}

type PendingReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RplTx [][]byte `protobuf:"bytes,1,rep,name=rplTx,proto3" json:"rplTx,omitempty"`
}

func (x *PendingReply) Reset() {
	*x = PendingReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_txpool_txpool_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PendingReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PendingReply) ProtoMessage() {}

func (x *PendingReply) ProtoReflect() protoreflect.Message {
	mi := &file_txpool_txpool_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PendingReply.ProtoReflect.Descriptor instead.
func (*PendingReply) Descriptor() ([]byte, []int) {
	return file_txpool_txpool_proto_rawDescGZIP(), []int{6}
}

func (x *PendingReply) GetRplTx() [][]byte {
	if x != nil {
		return x.RplTx
	}
	return nil
}

var File_txpool_txpool_proto protoreflect.FileDescriptor

var file_txpool_txpool_proto_rawDesc = []byte{
	0x0a, 0x13, 0x74, 0x78, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x74, 0x78, 0x70, 0x6f, 0x6f, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x74, 0x78, 0x70, 0x6f, 0x6f, 0x6c, 0x1a, 0x11, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x2f, 0x0a, 0x08, 0x54, 0x78, 0x48, 0x61, 0x73, 0x68, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x06,
	0x68, 0x61, 0x73, 0x68, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x48, 0x32, 0x35, 0x36, 0x52, 0x06, 0x68, 0x61, 0x73, 0x68, 0x65,
	0x73, 0x22, 0x21, 0x0a, 0x0d, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x78, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52,
	0x03, 0x74, 0x78, 0x73, 0x22, 0x3f, 0x0a, 0x0b, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x30, 0x0a, 0x08, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x74, 0x78, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x49,
	0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x08, 0x69, 0x6d, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x64, 0x22, 0x3d, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x23, 0x0a, 0x06, 0x68, 0x61, 0x73, 0x68, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x48, 0x32, 0x35, 0x36, 0x52, 0x06, 0x68, 0x61,
	0x73, 0x68, 0x65, 0x73, 0x22, 0x28, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x74, 0x78, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x03, 0x74, 0x78, 0x73, 0x22, 0x10,
	0x0a, 0x0e, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x24, 0x0a, 0x0c, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x72, 0x70, 0x6c, 0x54, 0x78, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52,
	0x05, 0x72, 0x70, 0x6c, 0x54, 0x78, 0x2a, 0x6c, 0x0a, 0x0c, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53,
	0x53, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x45,
	0x58, 0x49, 0x53, 0x54, 0x53, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x46, 0x45, 0x45, 0x5f, 0x54,
	0x4f, 0x4f, 0x5f, 0x4c, 0x4f, 0x57, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x54, 0x41, 0x4c,
	0x45, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x04,
	0x12, 0x12, 0x0a, 0x0e, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x10, 0x05, 0x32, 0x95, 0x02, 0x0a, 0x06, 0x54, 0x78, 0x70, 0x6f, 0x6f, 0x6c, 0x12,
	0x3d, 0x0a, 0x17, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x10, 0x2e, 0x74, 0x78, 0x70,
	0x6f, 0x6f, 0x6c, 0x2e, 0x54, 0x78, 0x48, 0x61, 0x73, 0x68, 0x65, 0x73, 0x1a, 0x10, 0x2e, 0x74,
	0x78, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x54, 0x78, 0x48, 0x61, 0x73, 0x68, 0x65, 0x73, 0x12, 0x40,
	0x0a, 0x12, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x15, 0x2e, 0x74, 0x78, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x49, 0x6d,
	0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x74, 0x78,
	0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x4f, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x1e, 0x2e, 0x74, 0x78, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x47, 0x65, 0x74,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x74, 0x78, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x47, 0x65, 0x74,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x39, 0x0a, 0x07, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x2e, 0x74,
	0x78, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x74, 0x78, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x50, 0x65,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x30, 0x01, 0x42, 0x11, 0x5a, 0x0f,
	0x2e, 0x2f, 0x74, 0x78, 0x70, 0x6f, 0x6f, 0x6c, 0x3b, 0x74, 0x78, 0x70, 0x6f, 0x6f, 0x6c, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_txpool_txpool_proto_rawDescOnce sync.Once
	file_txpool_txpool_proto_rawDescData = file_txpool_txpool_proto_rawDesc
)

func file_txpool_txpool_proto_rawDescGZIP() []byte {
	file_txpool_txpool_proto_rawDescOnce.Do(func() {
		file_txpool_txpool_proto_rawDescData = protoimpl.X.CompressGZIP(file_txpool_txpool_proto_rawDescData)
	})
	return file_txpool_txpool_proto_rawDescData
}

var file_txpool_txpool_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_txpool_txpool_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_txpool_txpool_proto_goTypes = []interface{}{
	(ImportResult)(0),              // 0: txpool.ImportResult
	(*TxHashes)(nil),               // 1: txpool.TxHashes
	(*ImportRequest)(nil),          // 2: txpool.ImportRequest
	(*ImportReply)(nil),            // 3: txpool.ImportReply
	(*GetTransactionsRequest)(nil), // 4: txpool.GetTransactionsRequest
	(*GetTransactionsReply)(nil),   // 5: txpool.GetTransactionsReply
	(*PendingRequest)(nil),         // 6: txpool.PendingRequest
	(*PendingReply)(nil),           // 7: txpool.PendingReply
	(*types.H256)(nil),             // 8: types.H256
}
var file_txpool_txpool_proto_depIdxs = []int32{
	8, // 0: txpool.TxHashes.hashes:type_name -> types.H256
	0, // 1: txpool.ImportReply.imported:type_name -> txpool.ImportResult
	8, // 2: txpool.GetTransactionsRequest.hashes:type_name -> types.H256
	1, // 3: txpool.Txpool.FindUnknownTransactions:input_type -> txpool.TxHashes
	2, // 4: txpool.Txpool.ImportTransactions:input_type -> txpool.ImportRequest
	4, // 5: txpool.Txpool.GetTransactions:input_type -> txpool.GetTransactionsRequest
	6, // 6: txpool.Txpool.Pending:input_type -> txpool.PendingRequest
	1, // 7: txpool.Txpool.FindUnknownTransactions:output_type -> txpool.TxHashes
	3, // 8: txpool.Txpool.ImportTransactions:output_type -> txpool.ImportReply
	5, // 9: txpool.Txpool.GetTransactions:output_type -> txpool.GetTransactionsReply
	7, // 10: txpool.Txpool.Pending:output_type -> txpool.PendingReply
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_txpool_txpool_proto_init() }
func file_txpool_txpool_proto_init() {
	if File_txpool_txpool_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_txpool_txpool_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxHashes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_txpool_txpool_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImportRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_txpool_txpool_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImportReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_txpool_txpool_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTransactionsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_txpool_txpool_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTransactionsReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_txpool_txpool_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PendingRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_txpool_txpool_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PendingReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_txpool_txpool_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_txpool_txpool_proto_goTypes,
		DependencyIndexes: file_txpool_txpool_proto_depIdxs,
		EnumInfos:         file_txpool_txpool_proto_enumTypes,
		MessageInfos:      file_txpool_txpool_proto_msgTypes,
	}.Build()
	File_txpool_txpool_proto = out.File
	file_txpool_txpool_proto_rawDesc = nil
	file_txpool_txpool_proto_goTypes = nil
	file_txpool_txpool_proto_depIdxs = nil
}