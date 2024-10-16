package grpc_adapter

import (
	"context"
	"log"
	"net"

	domain "github.com/MGomed/auth/internal/domain"
	api "github.com/MGomed/auth/pkg/user_api"

	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	reflection "google.golang.org/grpc/reflection"
)

// UserAPIUsecase interface of user_api usecase
type UserAPIUsecase interface {
	Create(ctx context.Context, req *domain.CreateRequest) (*domain.CreateResponse, error)
	Get(ctx context.Context, req *domain.GetRequest) (*domain.GetResponse, error)
	Update(ctx context.Context, req *domain.UpdateRequest) error
	Delete(ctx context.Context, req *domain.DeleteRequest) error
}

type GRPCConfig interface {
	Address() string
}

type server struct {
	api.UnimplementedUserAPIServer

	logger  *log.Logger
	config  GRPCConfig
	usecase UserAPIUsecase
}

// NewGrpcServer is server constructor
func NewGrpcServer(logger *log.Logger, config GRPCConfig, usecase UserAPIUsecase) *server {
	return &server{
		logger:  logger,
		config:  config,
		usecase: usecase,
	}
}

// Serve gets net.Listener and bind it to grpc server,
// also blocking execution by calling Serve()
func (s *server) Serve() error {
	lis, err := net.Listen("tcp", s.config.Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	reflection.Register(server)
	api.RegisterUserAPIServer(server, s)

	return server.Serve(lis)
}

// Create creates new user
func (s *server) Create(ctx context.Context, req *api.CreateRequest) (*api.CreateResponse, error) {
	resp, err := s.usecase.Create(ctx, domain.CreateReqFromAPIToDomain(req))
	if err != nil {
		return nil, err
	}

	return domain.CreateRespToAPIFromDomain(resp), nil
}

// Get gets user by id
func (s *server) Get(ctx context.Context, req *api.GetRequest) (*api.GetResponse, error) {
	resp, err := s.usecase.Get(ctx, domain.GetReqFromAPIToDomain(req))
	if err != nil {
		return nil, err
	}

	return domain.GetRespToAPIFromDomain(resp), nil
}

// Update modifies user information
func (s *server) Update(ctx context.Context, req *api.UpdateRequest) (*empty.Empty, error) {
	err := s.usecase.Update(ctx, domain.UpdateReqFromAPIToDomain(req))
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

// Delete removes user by id
func (s *server) Delete(ctx context.Context, req *api.DeleteRequest) (*empty.Empty, error) {
	err := s.usecase.Delete(ctx, domain.DeleteReqFromAPIToDomain(req))
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
