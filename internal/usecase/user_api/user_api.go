package userapi_usecase

import (
	"context"
	"log"

	domain "github.com/MGomed/auth/internal/domain"
)

type usecase struct {
	logger *log.Logger
}

// NewUserAPIUsecase is a usecase constructor
func NewUserAPIUsecase(logger *log.Logger) *usecase {
	return &usecase{
		logger: logger,
	}
}

// Create creates new user
func (uc *usecase) Create(_ context.Context, _ *domain.CreateRequest) (*domain.CreateResponse, error) {
	// TODO some business logic

	return nil, nil
}

// Get gets user by id
func (uc *usecase) Get(_ context.Context, _ *domain.GetRequest) (*domain.GetResponse, error) {
	// TODO some business logic

	return nil, nil
}

// Update modifies user information
func (uc *usecase) Update(_ context.Context, _ *domain.UpdateRequest) error {
	// TODO some business logic

	return nil
}

// Delete removes user by id
func (uc *usecase) Delete(_ context.Context, _ *domain.DeleteRequest) error {
	// TODO some business logic

	return nil
}
