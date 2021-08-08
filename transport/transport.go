package transport

import (
	"context"

	"github.com/GolangTechTask/pkg/api"
	apipb "github.com/GolangTechTask/pkg/api"
	"github.com/GolangTechTask/pkg/util"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
)

type grpcHandlers struct {
	CreateVoteAPI kitgrpc.Handler
	ListVoteAPI   kitgrpc.Handler
	CastVoteAPI   kitgrpc.Handler
}

func NewGRPC(e *Endpoints) apipb.VotingServiceServer {
	return grpcHandlers{
		CreateVoteAPI: kitgrpc.NewServer(
			e.CreateVote,
			util.NoReqResp,
			util.NoReqResp,
		),
		ListVoteAPI: kitgrpc.NewServer(
			e.ListVote,
			util.NoReqResp,
			util.NoReqResp,
		),
		CastVoteAPI: kitgrpc.NewServer(
			e.CastVote,
			util.NoReqResp,
			util.NoReqResp,
		),
	}
}

// DefaultHandlerCall validates the request and calls the given handler
func DefaultHandlerCall(ctx context.Context, api kitgrpc.Handler, req interface{}) (context.Context, interface{}, error) {
	return api.ServeGRPC(ctx, req)
}
func (g grpcHandlers) CreateVoteable(ctx context.Context, req *apipb.CreateVoteableRequest) (*apipb.CreateVoteableResponse, error) {
	_, resp, err := DefaultHandlerCall(ctx, g.CreateVoteAPI, req)
	if err != nil {
		return nil, err
	}
	return resp.(*api.CreateVoteableResponse), nil
}

func (g grpcHandlers) ListVoteables(ctx context.Context, req *apipb.ListVoteableRequest) (*apipb.ListVoteableResponse, error) {
	_, resp, err := DefaultHandlerCall(ctx, g.ListVoteAPI, req)
	if err != nil {
		return nil, err
	}
	return resp.(*api.ListVoteableResponse), nil
}

func (g grpcHandlers) CastVote(ctx context.Context, req *apipb.CastVoteRequest) (*apipb.CastVoteResponse, error) {
	_, resp, err := DefaultHandlerCall(ctx, g.CastVoteAPI, req)
	if err != nil {
		return nil, err
	}
	return resp.(*api.CastVoteResponse), nil
}
