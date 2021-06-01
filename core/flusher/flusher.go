package flusher

import (
	"fmt"

	"github.com/ozoncp/ocp-progress-api/core/progress"
	"github.com/ozoncp/ocp-progress-api/core/repo"
	"github.com/ozoncp/ocp-progress-api/internal/utils"
)

type Flusher interface {
	Flush(notes []progress.Pogress) []progress.Pogress
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

func (f *flusher) Flush(users []progress.Pogress) []progress.Pogress {

	chunks, err := utils.SplitToBulks(users, f.chunkSize)

	if err != nil {
		return users
	}

	for index, val := range chunks {
		fmt.Println("LOLOLO")
		if err := f.storage.AddProgress(val); err != nil {
			return users[index*f.chunkSize:]
		}
	}

	return nil
}
