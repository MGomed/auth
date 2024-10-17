package main

import (
	"context"
	"flag"
	"log"

	grpc_port "github.com/MGomed/auth/internal/adapter/grpc"
	postgres "github.com/MGomed/auth/internal/adapter/postgres"
	config "github.com/MGomed/auth/internal/config"
	env_config "github.com/MGomed/auth/internal/config/env"
	user_api "github.com/MGomed/auth/internal/usecase/user_api"
	logger "github.com/MGomed/auth/pkg/logger"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "build/.env", "path to config file")
	flag.Parse()
}

func main() {
	ctx := context.Background()

	err := env_config.Load(configPath)
	if err != nil {
		log.Fatal(err)
	}

	cfg := initConfig()

	logger, err := logger.InitLogger(cfg)
	if err != nil {
		log.Fatal(err)
	}

	pgAdapter, err := postgres.NewAdapter(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}

	userAPIUsecase := user_api.NewUserAPIUsecase(logger, pgAdapter)

	server := grpc_port.NewGrpcServer(logger, cfg, userAPIUsecase)

	logger.Println("Starting GRPC server!")

	if err := server.Serve(); err != nil {
		logger.Fatalf("failed to start grpc server: %v", err)
	}
}

func initConfig() *config.Config {
	grpcConfig, err := env_config.NewGRPCConfig()
	if err != nil {
		log.Fatal(err)
	}

	pgConfig, err := env_config.NewPgConfig()
	if err != nil {
		log.Fatal(err)
	}

	loggerConfig, err := env_config.NewLoggerConfig()
	if err != nil {
		log.Fatal(err)
	}

	return config.NewConfig(grpcConfig, pgConfig, loggerConfig)
}
