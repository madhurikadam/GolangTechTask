package service

import (
	"context"
	"reflect"
	"testing"

	apipb "github.com/GolangTechTask/pkg/api"
)

func Test_votingImpl_CastVote(t *testing.T) {
	type args struct {
		ctx     context.Context
		request *apipb.CastVoteRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *apipb.CastVoteResponse
		wantErr bool
	}{
		{
			name: "cast vote failed",
			args: args{
				ctx: context.Background(),
				request: &apipb.CastVoteRequest{
					Uuid: invalidUUID,
				},
			},
			wantErr: true,
			want: &apipb.CastVoteResponse{
				Success: false,
			},
		},
		{
			name: "cast vote succes",
			args: args{
				ctx: context.Background(),
				request: &apipb.CastVoteRequest{
					Uuid: "valid-uuid",
				},
			},
			wantErr: false,
			want: &apipb.CastVoteResponse{
				Success: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := provideMock()
			got, err := v.CastVote(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("votingImpl.CastVote() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("votingImpl.CreateVote() = %v, want %v", got, tt.want)
			}

		})
	}
}

func Test_votingImpl_CreateVote(t *testing.T) {
	type args struct {
		ctx     context.Context
		request *apipb.CreateVoteableRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "create vote fail",
			args: args{
				ctx: context.Background(),
				request: &apipb.CreateVoteableRequest{
					Question: invalidQuestion,
				},
			},
			wantErr: true,
		},
		{
			name: "create vote success",
			args: args{
				ctx: context.Background(),
				request: &apipb.CreateVoteableRequest{
					Question: "valid-question",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := provideMock()
			_, err := v.CreateVote(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("votingImpl.CreateVote() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_votingImpl_ListVote(t *testing.T) {
	type args struct {
		ctx     context.Context
		request *apipb.ListVoteableRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *apipb.ListVoteableResponse
		wantErr bool
	}{
		{
			name: "list vote fail",
			args: args{
				ctx: context.Background(),
				request: &apipb.ListVoteableRequest{
					NextPageToken: invalidUUID,
				},
			},
			wantErr: true,
		},
		{
			name: "list vote success",
			args: args{
				ctx: context.Background(),
				request: &apipb.ListVoteableRequest{
					NextPageToken: "valid-uuid",
				},
			},
			wantErr: false,
			want: &apipb.ListVoteableResponse{
				Votables: []*apipb.Voteable{
					{
						Uuid:     "valid-uuid",
						Question: "mock question",
						Answers:  []string{"mock1", "mock2"},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := provideMock()
			got, err := v.ListVote(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("votingImpl.ListVote() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("votingImpl.ListVote() = %v, want %v", got, tt.want)
			}
		})
	}
}
