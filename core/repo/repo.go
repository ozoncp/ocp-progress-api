package repo

import "github.com/ozoncp/ocp-progress-api/core/progress"

type Repo interface {
	AddProgress(progres []progress.Progress) error
	DescribeProgress(id uint64) (*progress.Progress, error)
	RemoveProgress(id uint64) error
	ListProgress(limit, offset uint64) ([]progress.Progress, error)
}
