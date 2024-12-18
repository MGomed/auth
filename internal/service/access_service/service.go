package accessservice

import (
	logger "github.com/MGomed/common/logger"
)

type service struct {
	log logger.Interface
}

// NewAccessService is a service struct constructor
func NewAccessService(log logger.Interface) *service {
	return &service{
		log: log,
	}
}
