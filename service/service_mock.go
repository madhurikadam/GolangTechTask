package service

import (
	"context"
	"errors"

	"github.com/GolangTechTask/model"
	"github.com/GolangTechTask/pkg/configuration"
	"github.com/GolangTechTask/pkg/constant"
	"github.com/GolangTechTask/pkg/logger"
	dbMock "github.com/GolangTechTask/repo/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

var (
	invalidUUID     = "invalid-uuid"
	invalidQuestion = "invalid-question"
)
var (
	errFailToCreate   = errors.New("fail to create")
	errFailToCastVote = errors.New("fail to cast vote")
	errFailToList     = errors.New("fail to list")
)

func provideMock() ServiceProvider {
	configuration.LoadDefaults()
	logger.Init(configuration.RequireInt(constant.LogLevel), configuration.RequireString(constant.LogTimeFormat))
	db := dbMock.VotingRepo{}
	db.On("CreateVote", mock.Anything, mock.Anything).Return(func(ctx context.Context, req *model.CreateVoteReq) *model.CreateVoteResp {
		if req.Question == invalidQuestion {
			return nil
		}
		return &model.CreateVoteResp{
			UUID: uuid.New().String(),
		}
	}, func(ctx context.Context, req *model.CreateVoteReq) error {
		if req.Question == invalidQuestion {
			return errFailToCreate
		}
		return nil
	})
	db.On("CastVote", mock.Anything, mock.Anything).Return(func(ctx context.Context, req *model.CastVoteReq) *model.CastVoteResp {
		if req.UUID == invalidUUID {
			return nil
		}
		return &model.CastVoteResp{
			UUID: req.UUID,
		}
	}, func(ctx context.Context, req *model.CastVoteReq) error {
		if req.UUID == invalidUUID {
			return errFailToCastVote
		}
		return nil
	})

	db.On("ListVote", mock.Anything, mock.Anything).Return(func(ctx context.Context, req *model.ListVoteReq) *model.ListVoteResp {
		if req.NextPageToken == invalidUUID {
			return nil
		}
		return &model.ListVoteResp{
			Resp: []model.VoteTable{
				{
					Vote:     1,
					Question: "mock question",
					UUID:     "valid-uuid",
					Answers:  []string{"mock1", "mock2"},
				},
			},
		}
	}, func(ctx context.Context, req *model.ListVoteReq) error {
		if req.NextPageToken == invalidUUID {
			return errFailToList
		}
		return nil
	})
	ser := votingImpl{
		db: &db,
	}

	return &ser
}
