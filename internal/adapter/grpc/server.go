package grpc_adapter

import (
	"context"
	"net"

	domain "github.com/MGomed/auth/internal/domain"
	api "github.com/MGomed/auth/pkg/user_api"

	gofakeit "github.com/brianvoe/gofakeit/v7"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	reflection "google.golang.org/grpc/reflection"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// UserAPIUsecase interface of user_api usecase
type UserAPIUsecase interface {
	Create(ctx context.Context, req *domain.CreateRequest) (*domain.CreateResponse, error)
	Get(ctx context.Context, req *domain.GetRequest) (*domain.GetResponse, error)
	Update(ctx context.Context, req *domain.UpdateRequest) error
	Delete(ctx context.Context, req *domain.DeleteRequest) error
}

type server struct {
	api.UnimplementedUserAPIServer

	usecase UserAPIUsecase
}

func NewGrpcServer(usecase UserAPIUsecase) *server {
	return &server{
		usecase: usecase,
	}
}

func (s *server) Serve(listener net.Listener) error {
	server := grpc.NewServer()
	reflection.Register(server)
	api.RegisterUserAPIServer(server, s)

	return server.Serve(listener)
}

func (s *server) Create(ctx context.Context, req *api.CreateRequest) (*api.CreateResponse, error) {
	//TODO correct the call ufter implementation
	_, _ = s.usecase.Create(ctx, domain.CreateReqFromAPIToDomain(req))

	return &api.CreateResponse{
		Id: gofakeit.Int64(),
	}, nil
}

func (s *server) Get(ctx context.Context, req *api.GetRequest) (*api.GetResponse, error) {
	// TODO correct the call ufter implementation
	_, _ = s.usecase.Get(ctx, domain.GetReqFromAPIToDomain(req))

	return &api.GetResponse{
		Id:        gofakeit.Int64(),
		Name:      gofakeit.Name(),
		Email:     gofakeit.Email(),
		Role:      api.Role_user,
		CreatedAt: timestamppb.Now(),
		UpdatedAt: timestamppb.Now(),
	}, nil
}

func (s *server) Update(ctx context.Context, req *api.UpdateRequest) (*empty.Empty, error) {
	// TODO correct the call ufter implementation
	_ = s.usecase.Update(ctx, domain.UpdateReqFromAPIToDomain(req))

	return &empty.Empty{}, nil
}

func (s *server) Delete(ctx context.Context, req *api.DeleteRequest) (*empty.Empty, error) {
	// TODO correct the call ufter implementation
	_ = s.usecase.Delete(ctx, domain.DeleteReqFromAPIToDomain(req))

	return &empty.Empty{}, nil
}
