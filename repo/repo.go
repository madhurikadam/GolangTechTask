package repo

import (
	"context"

	"github.com/GolangTechTask/model"
)

type VotingRepo interface {
	CreateVote(ctx context.Context, req *model.CreateVoteReq) (*model.CreateVoteResp, error)
	ListVote(ctx context.Context, req *model.ListVoteReq) (*model.ListVoteResp, error)
	CastVote(ctx context.Context, req *model.CastVoteReq) (*model.CastVoteResp, error)
}

func Init(ctx context.Context) (VotingRepo, error) {
	tableName := "voting"
	return InitDynamodb(ctx, tableName)
}
