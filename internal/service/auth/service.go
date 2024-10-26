package auth

import (
	"log"

	repository "github.com/MGomed/auth/internal/repository"
)

type service struct {
	logger *log.Logger
	repo   repository.Repository
}

// NewService is a service constructor
func NewService(logger *log.Logger, repo repository.Repository) *service {
	return &service{
		logger: logger,
		repo:   repo,
	}
}
