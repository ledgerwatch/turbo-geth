package stagedsync

import (
	"fmt"

	"github.com/ledgerwatch/turbo-geth/core"
	"github.com/ledgerwatch/turbo-geth/eth/stagedsync/stages"
	"github.com/ledgerwatch/turbo-geth/ethdb"
	"github.com/ledgerwatch/turbo-geth/log"
)

func DoStagedSyncWithFetchers(
	d DownloaderGlue,
	blockchain BlockChain,
	stateDB ethdb.Database,
	pid string,
	history bool,
	datadir string,
	quitCh chan struct{},
	headersFetchers []func() error,
) error {
	var err error
	defer log.Info("Staged sync finished")

	var syncHeadNumber uint64

	stages := []*Stage{
		{
			ID:          stages.Headers,
			Description: "Downloading headers",
			ExecFunc: func(s *StageState) error {
				return DownloadHeaders(s, d, stateDB, headersFetchers, quitCh)
			},
		},
		{
			ID:          stages.Bodies,
			Description: "Downloading block bodiess",
			ExecFunc: func(s *StageState) error {
				cont := true
				for cont && err == nil {
					fmt.Printf("cont=%v\n", cont)
					cont, err = spawnBodyDownloadStage(s, stateDB, d, pid)
					if err != nil {
						return err
					}
				}
				fmt.Printf("return s.Done\n")
				s.Done()
				return nil
			},
		},
		{
			ID:          stages.Senders,
			Description: "Recovering senders from tx signatures",
			ExecFunc: func(s *StageState) error {
				return spawnRecoverSendersStage(s, stateDB, blockchain.Config(), quitCh)
			},
		},
		{
			ID:          stages.Execution,
			Description: "Executing blocks w/o hash checks",
			ExecFunc: func(s *StageState) error {
				// TODO: Get rid of a global variable
				syncHeadNumber, err = spawnExecuteBlocksStage(s, stateDB, blockchain, quitCh)
				return err
			},
		},
		{
			ID:          stages.HashCheck,
			Description: "Validating final hash",
			ExecFunc: func(s *StageState) error {
				return spawnCheckFinalHashStage(s, stateDB, syncHeadNumber, datadir, quitCh)
			},
		},
		{
			ID:                  stages.AccountHistoryIndex,
			Description:         "Generating account history index",
			Disabled:            !history,
			DisabledDescription: "Enable by adding `h` to --storage-mode",
			ExecFunc: func(s *StageState) error {
				return spawnAccountHistoryIndex(s, stateDB, datadir, core.UsePlainStateExecution, quitCh)
			},
		},
		{
			ID:                  stages.StorageHistoryIndex,
			Description:         "Generating storage history index",
			Disabled:            !history,
			DisabledDescription: "Enable by adding `h` to --storage-mode",
			ExecFunc: func(s *StageState) error {
				return spawnStorageHistoryIndex(s, stateDB, datadir, core.UsePlainStateExecution, quitCh)
			},
		},
	}

	state := NewState(stages)

	i := 1

	for !state.IsDone() {
		stage := state.CurrentStage()

		if stage.Disabled {
			message := fmt.Sprintf(
				"Sync stage %d/%d. %v disabled. %s",
				i,
				state.Len(),
				stage.Description,
				stage.DisabledDescription,
			)

			log.Info(message)

			state.NextStage()
			i++
			continue
		}

		stageState, err := state.StageState(stage.ID, stateDB)
		if err != nil {
			return err
		}

		message := fmt.Sprintf("Sync stage %d/%d. %v...", i, state.Len(), stage.Description)
		log.Info(message)

		err = stage.ExecFunc(stageState)
		if err != nil {
			return err
		}

		log.Info(fmt.Sprintf("%s DONE!", message))

		i++
	}

	return nil
}
