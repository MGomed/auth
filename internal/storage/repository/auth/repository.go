package auth

import (
	db "github.com/MGomed/common/client/db"
)

type repository struct {
	dbc db.Client
}

// NewRepository is repository struct constructor
func NewRepository(dbc db.Client) *repository {
	return &repository{
		dbc: dbc,
	}
}
