package main

import (
	"context"
	"fmt"
	"net"

	apipb "github.com/GolangTechTask/pkg/api"
	"github.com/GolangTechTask/pkg/configuration"
	"github.com/GolangTechTask/pkg/constant"
	"github.com/GolangTechTask/pkg/logger"
	"github.com/GolangTechTask/pkg/middleware"
	"github.com/GolangTechTask/repo"

	"github.com/GolangTechTask/transport"
	"google.golang.org/grpc"

	"github.com/GolangTechTask/service"
)

func main() {
	ctx := context.Background()
	errors := make(chan error)

	go func() {
		configuration.LoadDefaults()
		if err := logger.Init(configuration.RequireInt(constant.LogLevel), configuration.RequireString(constant.LogTimeFormat)); err != nil {
			errors <- fmt.Errorf("failed to initialize logger: %v", err)
			return
		}
		port := configuration.RequireString(constant.LocalGrpcPort)
		listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
		if err != nil {
			errors <- err
			return
		}
		db, err := repo.Init(ctx)
		if err != nil {
			errors <- err
			return
		}
		s, err := service.New(ctx, db)
		if err != nil {
			errors <- err
			return
		}

		// gRPC server statup options
		opts := []grpc.ServerOption{}

		// add middleware
		opts = middleware.AddLogging(logger.Log, opts)

		e := transport.CreateEndpoints(s)
		grpcServer := transport.NewGRPC(e)
		// register service
		gRPCServer := grpc.NewServer(opts...)
		apipb.RegisterVotingServiceServer(gRPCServer, grpcServer)

		logger.Log.Info(fmt.Sprintf("gRPC listen on %s", port))
		errors <- gRPCServer.Serve(listener)
	}()
	err := <-errors
	logger.Log.Error(err.Error())
}
