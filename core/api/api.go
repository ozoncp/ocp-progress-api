package api

import (
	"context"

	"github.com/rs/zerolog/log"

	desc "github.com/ozoncp/ocp-progress-api/pkg/ocp-progress-api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	errProjectListEmpty = "not found any projects"
	errProjectNotFound  = "project not found"
	errProjectCreate    = "creating project fails"
	errProjectRemove    = "removing project fails"
)

type api struct {
	desc.UnimplementedOcpProgressApiServer
}

func NewOcpProgressApi() desc.OcpProgressApiServer {
	return &api{}
}

func (a *api) CreateProgressV1(
	ctx context.Context,
	req *desc.CreateProgressV1Request) (
	*desc.CreateProgressV1Response,
	error) {

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Info().Msgf("Got CreateProgressV1: {ClassroomId: %d, UserId: %d}", req.ClassroomId, req.UserId)

	err := status.Error(codes.NotFound, errProjectCreate)
	return nil, err
}

func (a *api) DescribeProgressV1(
	ctx context.Context,
	req *desc.DescribeProgressV1Request) (
	*desc.DescribeProgressV1Response,
	error) {

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Info().Msgf("Got DescribeProgressV1: {ProgressId: %d}", req.ProgressId)
	err := status.Error(codes.NotFound, errProjectCreate)
	return nil, err
}

func (a *api) ListProgressV1(
	ctx context.Context,
	req *desc.ListProgressV1Request) (
	*desc.ListProgressV1Response,
	error) {

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Info().Msgf("Got ListProgressV1: {limit: %d, offset: %d}", req.Limit, req.Offset)

	err := status.Error(codes.NotFound, errProjectListEmpty)
	return nil, err
}

func (a *api) RemoveProgressV1(
	ctx context.Context,
	req *desc.RemoveProgressV1Request) (
	*desc.RemoveProgressV1Response,
	error) {

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Info().Msgf("Got RemoveProgressV1: {ProgressId: %d}", req.ProgressId)

	err := status.Error(codes.NotFound, errProjectRemove)
	return nil, err

}
