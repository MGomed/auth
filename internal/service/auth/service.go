package auth

import (
	"log"

	repository "github.com/MGomed/auth/internal/repository"
	db "github.com/MGomed/auth/pkg/client/db"
)

type service struct {
	logger    *log.Logger
	repo      repository.Repository
	txManager db.TxManager
}

// NewService is a service constructor
func NewService(logger *log.Logger, repo repository.Repository, txManager db.TxManager) *service {
	return &service{
		logger:    logger,
		repo:      repo,
		txManager: txManager,
	}
}
