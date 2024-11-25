package userservice

import (
	"log"

	storage "github.com/MGomed/auth/internal/storage"
	db "github.com/MGomed/common/client/db"
)

type service struct {
	logger    *log.Logger
	repo      storage.Repository
	cache     storage.Cache
	txManager db.TxManager
	msgBus    storage.MessageBus
}

// NewUserService is a service constructor
func NewUserService(
	logger *log.Logger,
	repo storage.Repository,
	cache storage.Cache,
	txManager db.TxManager,
	msgBus storage.MessageBus,
) *service {
	return &service{
		logger:    logger,
		repo:      repo,
		cache:     cache,
		txManager: txManager,
		msgBus:    msgBus,
	}
}
