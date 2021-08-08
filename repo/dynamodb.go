package repo

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/GolangTechTask/model"
	"github.com/GolangTechTask/pkg/configuration"
	"github.com/GolangTechTask/pkg/constant"
	"github.com/GolangTechTask/pkg/logger"
	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

type DynamoDB struct {
	table *dynamo.Table
}

func InitDynamodb(ctx context.Context, tableName string) (*DynamoDB, error) {
	endpoint := configuration.RequireString(constant.DbEndPoint)
	region := configuration.RequireString(constant.AWSRegion)
	key := configuration.RequireString(constant.AWSKey)
	secret := configuration.RequireString(constant.AWSSecret)
	config := aws.NewConfig().WithEndpoint(endpoint).WithRegion(region).WithCredentials(credentials.NewStaticCredentials(key, secret, ""))
	db := dynamo.New(session.New(), config)
	err := db.CreateTable(tableName, model.VoteTable{}).Run()
	if err != nil {
		if err.(awserr.Error).Code() == "ResourceInUseException" {
			logger.Log.Info("table is already exist ")
		} else {
			logger.Log.Error(fmt.Sprintf("error while creating table %v", err))
			return nil, err
		}
	}
	table := db.Table(tableName)
	return &DynamoDB{
		table: &table,
	}, nil
}

func (db *DynamoDB) CreateVote(ctx context.Context, req *model.CreateVoteReq) (*model.CreateVoteResp, error) {
	err := db.table.Put(model.VoteTable{
		UUID:     req.UUID,
		Question: req.Question,
		Answers:  req.Answers,
	}).RunWithContext(ctx)
	if err != nil {
		return nil, err
	}
	return &model.CreateVoteResp{
		UUID: req.UUID,
	}, nil
}

func (db *DynamoDB) ListVote(ctx context.Context, req *model.ListVoteReq) (*model.ListVoteResp, error) {
	var results []model.VoteTable
	sc := db.table.Scan()
	if req.PageSize > 0 {
		sc = sc.Limit(req.PageSize)
	}
	if req.NextPageToken != "" {
		pkBytes, err := base64.StdEncoding.DecodeString(req.NextPageToken)
		if err != nil {
			return nil, err
		}
		var pageKey dynamo.PagingKey
		if err = json.Unmarshal(pkBytes, &pageKey); err != nil {
			return nil, err
		}
		sc = sc.StartFrom(pageKey)
	}

	nextPageKey, err := sc.AllWithLastEvaluatedKey(&results)
	if err != nil {
		return nil, err
	}
	nextPageBytes, err := json.Marshal(nextPageKey)
	if err != nil {
		return nil, err
	}
	nextPage := base64.StdEncoding.EncodeToString(nextPageBytes)
	resp := &model.ListVoteResp{
		Resp:          results,
		NextPageToken: nextPage,
	}
	return resp, nil
}

func (db *DynamoDB) CastVote(ctx context.Context, req *model.CastVoteReq) (*model.CastVoteResp, error) {
	_, err := db.validate(req)
	if err != nil {
		return nil, err
	}
	err = db.table.Update("ID", req.UUID).Set("Vote", req.VoteIndex).Run()
	if err != nil {
		return nil, err
	}
	return &model.CastVoteResp{UUID: req.UUID}, nil
}

func (db *DynamoDB) validate(req *model.CastVoteReq) (bool, error) {

	var result *model.VoteTable
	err := db.table.Get("ID", req.UUID).One(&result)
	if err != nil {
		return false, err
	}
	return true, nil
}
