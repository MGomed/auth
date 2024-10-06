package main

import (
	"fmt"
	"log"
	"net"

	config "github.com/MGomed/auth/config"
	grpc_adapter "github.com/MGomed/auth/internal/adapter/grpc"
	user_api "github.com/MGomed/auth/internal/usecase/user_api"
	logger "github.com/MGomed/auth/pkg/logger.go"
)

func main() {
	conf, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	log, err := logger.InitLogger(conf)
	if err != nil {
		log.Fatal(err)
	}

	userAPIUsecase := user_api.NewUserAPIUsecase(log)

	server := grpc_adapter.NewGrpcServer(log, userAPIUsecase)

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", conf.GRPC.Host, conf.GRPC.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("Starting GRPC server!")

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to start grpc server: %v", err)
	}
}
