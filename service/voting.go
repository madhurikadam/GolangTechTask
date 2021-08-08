package service

import (
	"context"

	"github.com/GolangTechTask/model"
	apipb "github.com/GolangTechTask/pkg/api"
	"github.com/GolangTechTask/pkg/logger"
	"github.com/GolangTechTask/repo"
	"github.com/google/uuid"
)

type votingImpl struct {
	db repo.VotingRepo
}

// CreateVote create the voting question
func (v *votingImpl) CreateVote(ctx context.Context, request *apipb.CreateVoteableRequest) (*apipb.CreateVoteableResponse, error) {
	logger.Log.Info("Create Vote Question")
	uuid := uuid.New().String()
	resp, err := v.db.CreateVote(ctx, &model.CreateVoteReq{
		UUID:     uuid,
		Question: request.Question,
		Answers:  request.Answers,
	})
	if err != nil {
		logger.Log.Error(err.Error())
		return nil, err
	}
	return &apipb.CreateVoteableResponse{
		Uuid: resp.UUID,
	}, nil
}

//ListVote list all voting question, it support the pagination user pageSize and nextPageToken to manage pagination
func (v *votingImpl) ListVote(ctx context.Context, request *apipb.ListVoteableRequest) (*apipb.ListVoteableResponse, error) {
	logger.Log.Info("List Votes")
	resp, err := v.db.ListVote(ctx, &model.ListVoteReq{
		PageSize:      request.GetPageSize(),
		NextPageToken: request.GetNextPageToken(),
	})
	if err != nil {
		logger.Log.Error(err.Error())
		return nil, err
	}
	listResp := []*apipb.Voteable{}
	for _, item := range resp.Resp {
		re := &apipb.Voteable{
			Uuid:     item.UUID,
			Question: item.Question,
			Answers:  item.Answers,
		}
		listResp = append(listResp, re)
	}
	return &apipb.ListVoteableResponse{
		Votables:      listResp,
		NextPageToken: resp.NextPageToken,
	}, nil
}

// CastVote cast the answer for give question
func (v *votingImpl) CastVote(ctx context.Context, request *apipb.CastVoteRequest) (*apipb.CastVoteResponse, error) {
	logger.Log.Info("Cast Vote")
	_, err := v.db.CastVote(ctx, &model.CastVoteReq{UUID: request.GetUuid(), VoteIndex: request.GetAnswerIndex()})
	if err != nil {
		logger.Log.Error(err.Error())
		return &apipb.CastVoteResponse{
			Success: false,
		}, err
	}
	return &apipb.CastVoteResponse{
		Success: true,
	}, nil
}
