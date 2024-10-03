package main

import (
	"fmt"
	"log"
	"net"

	grpc_adapter "github.com/MGomed/auth/internal/adapter/grpc"
	user_api "github.com/MGomed/auth/internal/usecase/user_api"
)

const (
	grpcPort = 50051
)

func main() {
	userAPIUsecase := user_api.NewUserAPIUsecase()

	server := grpc_adapter.NewGrpcServer(userAPIUsecase)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to start grpc server: %v", err)
	}
}
