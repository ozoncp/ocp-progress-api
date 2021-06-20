package api

import (
	"context"
	"errors"

	"github.com/rs/zerolog/log"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/ozoncp/ocp-progress-api/core/flusher"
	"github.com/ozoncp/ocp-progress-api/core/progress"
	"github.com/ozoncp/ocp-progress-api/core/repo"
	"github.com/ozoncp/ocp-progress-api/internal/metrics"
	"github.com/ozoncp/ocp-progress-api/internal/producer"
	"github.com/ozoncp/ocp-progress-api/internal/utils"
	desc "github.com/ozoncp/ocp-progress-api/pkg/ocp-progress-api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//const (
//	errProjectListEmpty = "not found any projects"
//	errProjectCreate    = "creating project fails"
//	errProjectRemove    = "removing project fails"
//)

const chunkSize int = 10

type api struct {
	desc.UnimplementedOcpProgressApiServer
	progressRepo repo.Repo
	logProducer  producer.LogProducer
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

	return &desc.DescribeProgressV1Response{Progress: progress.ToProtoProgress()}, nil
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
		protoProgress = append(protoProgress, progress.ToProtoProgress())
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

func (a *api) MultiCreateProgressV1(
	ctx context.Context,
	req *desc.MultiCreateProgressV1Request) (
	res *desc.MultiCreateProgressV1Response,
	err error) {

	defer utils.LogGrpcCall("MultiCreateProgressV1", &req, &res, &err)

	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("MultiCreateProgressV1")
	defer span.Finish()

	if err = req.Validate(); err != nil {

		err = status.Error(codes.InvalidArgument, err.Error())
		return nil, err
	}

	var progressSlice []progress.Progress
	for _, protoProgress := range req.Progress {

		progressSlice = append(progressSlice, progress.Progress{
			Id:             0,
			ClassroomId:    uint(protoProgress.ClassroomId),
			PresentationId: uint(protoProgress.PresentationId),
			SlideId:        uint(protoProgress.SlideId),
			UserId:         uint(protoProgress.UserId),
		})
	}

	fl := flusher.New(a.progressRepo, chunkSize)
	remainingProgress := fl.Flush(ctx, span, progressSlice)

	var createdCount = uint64(len(progressSlice) - len(remainingProgress))
	if createdCount == 0 {

		err = status.Error(codes.Unavailable, errors.New("flush call returned non nil result").Error())
		return nil, err
	}

	res = &desc.MultiCreateProgressV1Response{NumberOfProgressCreated: createdCount}
	return res, nil
}

func (a *api) UpdateProgressroomV1(
	ctx context.Context,
	req *desc.UpdateProgressV1Request) (
	res *desc.UpdateProgressV1Response,
	err error) {

	defer utils.LogGrpcCall("UpdateProgressroomV1", &req, &res, &err)
	defer func() {
		_ = a.logProducer.Send(producer.Updated, req, res, err)
	}()

	if err = req.Validate(); err != nil {

		err = status.Error(codes.InvalidArgument, err.Error())
		return nil, err
	}

	progressFrom := progress.FromProtoProgress(req.Note)

	found, err := a.progressRepo.UpdateClassroom(ctx, *progressFrom)
	if err != nil {

		err = status.Error(codes.Unavailable, err.Error())
		return nil, err
	}

	if bool(found) {
		metrics.IncUpdateCounter()
	}

	res = &desc.UpdateProgressV1Response{Found: found}
	return res, nil
}
