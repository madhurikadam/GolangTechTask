package service

import (
	"context"
	"sync"

	apipb "github.com/GolangTechTask/pkg/api"
	"github.com/GolangTechTask/repo"
)

//ServiceProvider can inclue multiple services
type ServiceProvider interface {
	VotingService
}

//VotingService ... Service for voting app
type VotingService interface {
	CreateVote(ctx context.Context, request *apipb.CreateVoteableRequest) (*apipb.CreateVoteableResponse, error)
	ListVote(ctx context.Context, request *apipb.ListVoteableRequest) (*apipb.ListVoteableResponse, error)
	CastVote(ctx context.Context, request *apipb.CastVoteRequest) (*apipb.CastVoteResponse, error)
}

// New initialize and configure the service instance
func New(ctx context.Context, db repo.VotingRepo) (ServiceProvider, error) {
	votingObj, err := NewVoting(ctx, db)
	if err != nil {
		return nil, err
	}
	return &struct {
		VotingService
	}{
		votingObj,
	}, err
}

var initVotingSvc sync.Once
var votingSvc VotingService

//NewVoting create/return voting service instance
func NewVoting(ctx context.Context, db repo.VotingRepo) (VotingService, error) {
	initVotingSvc.Do(func() {
		votingSvc = &votingImpl{
			db: db,
		}
	})
	return votingSvc, nil
}
