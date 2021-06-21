package flusher

import (
	"context"

	"github.com/ozoncp/ocp-progress-api/core/progress"
	"github.com/ozoncp/ocp-progress-api/core/repo"
	"github.com/ozoncp/ocp-progress-api/internal/utils"
)

type Flusher interface {
	Flush(ctx context.Context, notes []progress.Progress) []progress.Progress
}

type flusher struct {
	chunkSize int
	storage   repo.Repo
}

func New(storage repo.Repo, chSize int) Flusher {
	return &flusher{
		chunkSize: chSize,
		storage:   storage,
	}
}

func (f *flusher) Flush(ctx context.Context, progressSlice []progress.Progress) []progress.Progress {

	chunks, err := utils.SplitToBulks(progressSlice, f.chunkSize)

	if err != nil {
		return progressSlice
	}

	for index, val := range chunks {

		if err := f.storage.AddProgress(ctx, val); err != nil {
			return progressSlice[index*f.chunkSize:]
		}
	}

	return nil
}
