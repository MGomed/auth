package authservice

import (
	"log"

	storage "github.com/MGomed/auth/internal/storage"
)

type service struct {
	log  *log.Logger
	repo storage.Repository
}

// NewAuthService is a service struct constructor
func NewAuthService(log *log.Logger, repo storage.Repository) *service {
	return &service{
		log:  log,
		repo: repo,
	}
}
