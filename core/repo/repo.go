package repo

import "github.com/ozoncp/ocp-progress-api/core/progress"

type Repo interface {
	AddProgress(progres []progress.Pogress) error
	DescribeProgress(id uint64) (*progress.Pogress, error)
	RemoveProgress(id uint64) error
	ListProgress(limit, offset uint64) ([]progress.Pogress, error)
}
