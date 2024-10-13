package userapi_usecase

import (
	"context"
	"log"

	domain "github.com/MGomed/auth/internal/domain"
)

type DatabaseAdapter interface {
	CreateUser(ctx context.Context, info *domain.UserInfo) (int, error)
	GetUser(ctx context.Context, id int) (*domain.UserInfo, error)
	UpdateUser(ctx context.Context, info *domain.UserInfo) (int, error)
	DeleteUser(ctx context.Context, id int) (int, error)
}

type usecase struct {
	logger    *log.Logger
	dbAdapter DatabaseAdapter
}

// NewUserAPIUsecase is a usecase constructor
func NewUserAPIUsecase(logger *log.Logger, dbAdapter DatabaseAdapter) *usecase {
	return &usecase{
		logger:    logger,
		dbAdapter: dbAdapter,
	}
}

// Create creates new user
func (uc *usecase) Create(ctx context.Context, req *domain.CreateRequest) (*domain.CreateResponse, error) {
	id, err := uc.dbAdapter.CreateUser(ctx, req.Info)
	if err != nil {
		uc.logger.Printf("Failed to add user %v in database: %v", req.Info, err)

		return nil, err
	}

	uc.logger.Printf("Successfully added user with id: %v", id)

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
