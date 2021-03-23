package stagedsync

import "github.com/ledgerwatch/turbo-geth/turbo/stages/bodydownload"

type DownloaderGlue interface {
	SpawnHeaderDownloadStage([]func() error, *StageState, Unwinder) error
	SpawnBodyDownloadStage(string, string, *StageState, Unwinder, *bodydownload.PrefetchedBlocks) (bool, error)
}
