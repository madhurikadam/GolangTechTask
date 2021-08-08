package model

type CreateVoteReq struct {
	UUID     string `dynamo:"ID,hash"`
	Question string
	Answers  []string
}

type CreateVoteResp struct {
	UUID string `dynamo:"ID,hash"`
}

type ListVoteReq struct {
	PageSize      int64
	NextPageToken string
}

type ListVoteResp struct {
	Resp          []VoteTable
	NextPageToken string
}

type CastVoteResp struct {
	UUID string `dynamo:"ID,hash"`
}
type CastVoteReq struct {
	UUID      string
	VoteIndex int64
}

type VoteTable struct {
	Vote     int64
	UUID     string `dynamo:"ID,hash"`
	Question string
	Answers  []string
}
