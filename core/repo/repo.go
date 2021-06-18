package repo

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"

	"github.com/ozoncp/ocp-progress-api/core/progress"
)

const tableName = "progress"

type progressRepo struct {
	db *sql.DB
}

func New(db *sql.DB) Repo {

	return &progressRepo{db: db}
}

type Repo interface {
	AddProgress(ctx context.Context, progress []progress.Progress) error
	AddOneProgress(ctx context.Context, progress progress.Progress) (uint64, error)
	DescribeProgress(ctx context.Context, id uint64) (*progress.Progress, error)
	RemoveProgress(ctx context.Context, id uint64) error
	ListProgress(ctx context.Context, limit, offset uint64) ([]progress.Progress, error)
}

func (pr *progressRepo) AddProgress(ctx context.Context, progress []progress.Progress) error {
	query := sq.Insert(tableName).
		Columns("classroom_id", "presentation_id", "slide_id", "user_id").
		RunWith(pr.db).
		PlaceholderFormat(sq.Dollar)

	for _, pr := range progress {
		query = query.Values(pr.ClassroomId, pr.PresentationId, pr.SlideId, pr.UserId)
	}
	_, err := query.ExecContext(ctx)

	return err
}

func (pr *progressRepo) AddOneProgress(ctx context.Context, progress progress.Progress) (uint64, error) {

	query := sq.Insert(tableName).
		Columns("classroom_id", "presentation_id", "slide_id", "user_id").
		RunWith(pr.db).
		PlaceholderFormat(sq.Dollar)

	err := query.QueryRowContext(ctx).Scan(&progress.Id)
	if err != nil {
		return 0, err
	}

	return progress.Id, nil
}

func (pr *progressRepo) DescribeProgress(ctx context.Context, id uint64) (*progress.Progress, error) {
	query := sq.Select("id", "classroom_id", "presentation_id", "slide_id", "user_id").
		From(tableName).
		Where(sq.Eq{"id": id}).
		RunWith(pr.db).
		PlaceholderFormat(sq.Dollar)

	var progress progress.Progress

	if err := query.QueryRowContext(ctx).Scan(&progress.Id, &progress.ClassroomId, &progress.PresentationId, &progress.SlideId, &progress.UserId); err != nil {
		return nil, err
	}
	return &progress, nil
}

func (pr *progressRepo) RemoveProgress(ctx context.Context, id uint64) error {
	query := sq.Delete(tableName).
		Where(sq.Eq{"id": id}).
		RunWith(pr.db).
		PlaceholderFormat(sq.Dollar)

	_, err := query.ExecContext(ctx)

	if err != nil {
		return err
	}

	return nil
}

func (pr *progressRepo) ListProgress(ctx context.Context, limit, offset uint64) ([]progress.Progress, error) {

	query := sq.Select("id", "classroom_id", "presentation_id", "slide_id", "user_id").
		From(tableName).
		RunWith(pr.db).
		Limit(limit).
		Offset(offset).
		PlaceholderFormat(sq.Dollar)

	var progressSlice []progress.Progress

	rows, err := query.QueryContext(ctx)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var progress progress.Progress
		err = rows.Scan(&progress.Id, &progress.ClassroomId, &progress.PresentationId, &progress.SlideId, &progress.UserId)

		if err != nil {
			return nil, err
		}

		progressSlice = append(progressSlice, progress)
	}

	return progressSlice, nil
}
