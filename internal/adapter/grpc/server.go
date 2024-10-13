package grpc_adapter

import (
	"context"
	"log"
	"net"

	domain "github.com/MGomed/auth/internal/domain"
	api "github.com/MGomed/auth/pkg/user_api"

	gofakeit "github.com/brianvoe/gofakeit/v7"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	reflection "google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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
	opt := protojson.MarshalOptions{Indent: "    "} // for beautiful logs
	msg, _ := opt.Marshal(req)
	s.logger.Printf("<<<< Received create request:\n%s", msg)

	//TODO correct the call ufter implementation
	_, _ = s.usecase.Create(ctx, domain.CreateReqFromAPIToDomain(req))

	resp := &api.CreateResponse{
		Id: gofakeit.Int64(),
	}

	msg, _ = opt.Marshal(resp)
	s.logger.Printf(">>>> Sent create response:\n%s", msg)

	return resp, nil
}

// Get gets user by id
func (s *server) Get(ctx context.Context, req *api.GetRequest) (*api.GetResponse, error) {
	opt := protojson.MarshalOptions{Indent: "    "} // for beautiful logs
	msg, _ := opt.Marshal(req)
	s.logger.Printf("<<<< Received get request:\n%s", msg)

	// TODO correct the call ufter implementation
	_, _ = s.usecase.Get(ctx, domain.GetReqFromAPIToDomain(req))

	resp := &api.GetResponse{
		Id:        gofakeit.Int64(),
		Name:      gofakeit.Name(),
		Email:     gofakeit.Email(),
		Role:      api.Role_USER,
		CreatedAt: timestamppb.Now(),
		UpdatedAt: timestamppb.Now(),
	}

	msg, _ = opt.Marshal(resp)
	s.logger.Printf(">>>> Sent get response:\n%s", msg)

	return resp, nil
}

// Update modifies user information
func (s *server) Update(ctx context.Context, req *api.UpdateRequest) (*empty.Empty, error) {
	opt := protojson.MarshalOptions{Indent: "    "} // for beautiful logs
	msg, _ := opt.Marshal(req)
	s.logger.Printf("<<<< Received update request:\n%s", msg)

	// TODO correct the call ufter implementation
	_ = s.usecase.Update(ctx, domain.UpdateReqFromAPIToDomain(req))

	return &empty.Empty{}, nil
}

// Delete removes user by id
func (s *server) Delete(ctx context.Context, req *api.DeleteRequest) (*empty.Empty, error) {
	opt := protojson.MarshalOptions{Indent: "    "} // for beautiful logs
	msg, _ := opt.Marshal(req)
	s.logger.Printf("<<<< Received delete request:\n%s", msg)

	// TODO correct the call ufter implementation
	_ = s.usecase.Delete(ctx, domain.DeleteReqFromAPIToDomain(req))

	return &empty.Empty{}, nil
}
