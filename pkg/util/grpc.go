package util

import (
	"context"
)

// NoReqResp is a decoder
func NoReqResp(ctx context.Context, req interface{}) (interface{}, error) {
	return req, nil
}
