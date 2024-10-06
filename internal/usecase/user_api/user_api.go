package userapi_usecase

import (
	"context"
	"log"

	domain "github.com/MGomed/auth/internal/domain"
)

type usecase struct {
	logger *log.Logger
}

func NewUserAPIUsecase(logger *log.Logger) *usecase {
	return &usecase{
		logger: logger,
	}
}

func (uc *usecase) Create(ctx context.Context, req *domain.CreateRequest) (*domain.CreateResponse, error) {
	// TODO some business logic

	return nil, nil
}

func (uc *usecase) Get(ctx context.Context, req *domain.GetRequest) (*domain.GetResponse, error) {
	// TODO some business logic

	return nil, nil
}

func (uc *usecase) Update(ctx context.Context, req *domain.UpdateRequest) error {
	// TODO some business logic

	return nil
}

func (uc *usecase) Delete(ctx context.Context, req *domain.DeleteRequest) error {
	// TODO some business logic

	return nil
}
