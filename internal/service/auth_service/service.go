package authservice

import (
	storage "github.com/MGomed/auth/internal/storage"
	logger "github.com/MGomed/common/logger"
)

type service struct {
	log  logger.Interface
	repo storage.Repository
}

// NewAuthService is a service struct constructor
func NewAuthService(log logger.Interface, repo storage.Repository) *service {
	return &service{
		log:  log,
		repo: repo,
	}
}
