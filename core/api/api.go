package api

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/ozoncp/ocp-progress-api/core/progress"
	"github.com/ozoncp/ocp-progress-api/core/repo"
	desc "github.com/ozoncp/ocp-progress-api/pkg/ocp-progress-api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	errProjectListEmpty = "not found any projects"
	errProjectCreate    = "creating project fails"
	errProjectRemove    = "removing project fails"
)

type api struct {
	desc.UnimplementedOcpProgressApiServer
	progressRepo repo.Repo
}

func NewOcpProgressApi(progressRepo repo.Repo) desc.OcpProgressApiServer {
	return &api{
		UnimplementedOcpProgressApiServer: desc.UnimplementedOcpProgressApiServer{},
		progressRepo:                      progressRepo,
	}
}

func (a *api) CreateProgressV1(
	ctx context.Context,
	req *desc.CreateProgressV1Request) (
	*desc.CreateProgressV1Response,
	error) {

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	id, err := a.progressRepo.AddOneProgress(ctx, progress.Progress{
		Id:             0,
		ClassroomId:    uint(req.ClassroomId),
		PresentationId: uint(req.PresentationId),
		SlideId:        uint(req.SlideId),
		UserId:         uint(req.UserId),
	})

	if err != nil {
		log.Error().
			Err(err).
			Msg("Failed to add progress")
		return nil, err
	}

	log.Info().
		Uint64("ClassroomId", req.ClassroomId).
		Uint64("UserId", req.UserId).
		Msg("Got CreateProgressV1")

	//err := status.Error(codes.NotFound, errProjectCreate)
	return &desc.CreateProgressV1Response{Id: id}, nil
}

func (a *api) DescribeProgressV1(
	ctx context.Context,
	req *desc.DescribeProgressV1Request) (
	*desc.DescribeProgressV1Response,
	error) {

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	progress, err := a.progressRepo.DescribeProgress(ctx, req.ProgressId)

	if err != nil {
		log.Error().
			Err(err).
			Msg("Failed to describe progress")
		return nil, err
	}

	log.Info().
		Uint64("ProgressId", req.ProgressId).
		Msg("Got DescribeProgressV1")

	//err := status.Error(codes.NotFound, errProjectCreate)

	return &desc.DescribeProgressV1Response{Progress: progress.ToProtoClassroom()}, nil
}

func (a *api) ListProgressV1(
	ctx context.Context,
	req *desc.ListProgressV1Request) (
	*desc.ListProgressV1Response,
	error) {

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	progress, err := a.progressRepo.ListProgress(ctx, req.Limit, req.Offset)

	if err != nil {
		log.Error().
			Err(err).
			Msg("Failed to List progress")
		return nil, err
	}

	var protoProgress []*desc.Progress

	for _, progress := range progress {
		protoProgress = append(protoProgress, progress.ToProtoClassroom())
	}

	log.Info().
		Uint64("Limit", req.Limit).
		Uint64("Offset", req.Offset).
		Msg("Got ListProgressV1")

	//err := status.Error(codes.NotFound, errProjectListEmpty)
	return &desc.ListProgressV1Response{Progress: protoProgress}, nil
}

func (a *api) RemoveProgressV1(
	ctx context.Context,
	req *desc.RemoveProgressV1Request) (
	*desc.RemoveProgressV1Response,
	error) {

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err := a.progressRepo.RemoveProgress(ctx, req.ProgressId)

	if err != nil {
		log.Error().
			Err(err).
			Uint64("ProgressId", req.ProgressId).
			Msg("Failed to remove progress")
		return &desc.RemoveProgressV1Response{HasRemoved: false}, err
	}
	log.Info().
		Uint64("ProgressId", req.ProgressId).
		Msg("Got RemoveProgressV1")

	//err := status.Error(codes.NotFound, errProjectRemove)
	return &desc.RemoveProgressV1Response{HasRemoved: true}, nil

}
