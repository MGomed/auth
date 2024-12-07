package accessservice

import "log"

type service struct {
	log *log.Logger
}

// NewAccessService is a service struct constructor
func NewAccessService(log *log.Logger) *service {
	return &service{
		log: log,
	}
}
