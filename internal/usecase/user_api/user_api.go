package userapi_usecase

import (
	"context"
	"log"

	domain "github.com/MGomed/auth/internal/domain"
)

// DatabaseAdapter declaired interface for database communication
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

	return &domain.CreateResponse{
		ID: int64(id),
	}, nil
}

// Get gets user by id
func (uc *usecase) Get(ctx context.Context, req *domain.GetRequest) (*domain.GetResponse, error) {
	info, err := uc.dbAdapter.GetUser(ctx, int(req.ID))
	if err != nil {
		uc.logger.Printf("Failed to get user with id - %v from database: %v", req.ID, err)

		return nil, err
	}

	uc.logger.Printf("Successfully got user: %v", info)

	return &domain.GetResponse{
		Info: info,
	}, nil
}

// Update modifies user information
func (uc *usecase) Update(ctx context.Context, req *domain.UpdateRequest) error {
	_, err := uc.dbAdapter.UpdateUser(ctx, req.Info)
	if err != nil {
		uc.logger.Printf("Failed to update user - %v from database: %v", req.Info, err)

		return err
	}

	uc.logger.Printf("Successfully updated user with id: %v", req.Info.ID)

	return nil
}

// Delete removes user by id
func (uc *usecase) Delete(ctx context.Context, req *domain.DeleteRequest) error {
	_, err := uc.dbAdapter.DeleteUser(ctx, int(req.ID))
	if err != nil {
		uc.logger.Printf("Failed to delete user with id - %v from database: %v", req.ID, err)

		return err
	}

	uc.logger.Printf("Successfully deleted user: %v", req.ID)

	return nil
}
