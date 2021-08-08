package transport

import (
	"context"

	"github.com/GolangTechTask/pkg/api"
	"github.com/GolangTechTask/pkg/logger"
	"github.com/GolangTechTask/service"
	"github.com/go-kit/kit/endpoint"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	//ErrInvalidReq invalid request
	ErrInvalidReq = status.Errorf(codes.InvalidArgument, "Invalid Request")
)

type Endpoints struct {
	CreateVote endpoint.Endpoint
	ListVote   endpoint.Endpoint
	CastVote   endpoint.Endpoint
}

func CreateEndpoints(service service.ServiceProvider) *Endpoints {
	return &Endpoints{
		CreateVote: createVoteEP(service),
		ListVote:   listVoteEP(service),
		CastVote:   castVoteEP(service),
	}
}

func createVoteEP(service service.ServiceProvider) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(*api.CreateVoteableRequest)
		if !ok {
			logger.Log.Error(ErrInvalidReq.Error())
			return nil, ErrInvalidReq
		}
		return service.CreateVote(ctx, req)
	}
}

func listVoteEP(service service.ServiceProvider) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(*api.ListVoteableRequest)
		if !ok {
			logger.Log.Error(ErrInvalidReq.Error())
			return nil, ErrInvalidReq
		}
		return service.ListVote(ctx, req)
	}
}

func castVoteEP(service service.ServiceProvider) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(*api.CastVoteRequest)
		if !ok {
			logger.Log.Error(ErrInvalidReq.Error())
			return nil, ErrInvalidReq
		}
		return service.CastVote(ctx, req)
	}
}
