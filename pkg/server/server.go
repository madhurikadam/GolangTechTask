package server

import (
	"google.golang.org/grpc"
)

type Server struct {
	GRPCServer *grpc.Server
}
