package userservice

import (
	storage "github.com/MGomed/auth/internal/storage"
	db "github.com/MGomed/common/client/db"
	logger "github.com/MGomed/common/logger"
)

type service struct {
	logger    logger.Interface
	repo      storage.Repository
	cache     storage.Cache
	txManager db.TxManager
	msgBus    storage.MessageBus
}

// NewUserService is a service constructor
func NewUserService(
	logger logger.Interface,
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
