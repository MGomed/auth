package auth

import (
	"errors"

	db "github.com/MGomed/auth/pkg/client/db"
)

const (
	authTable = "auth"

	idColumn        = "id"
	nameColumn      = "name"
	emailColumn     = "email"
	passwordColumn  = "password"
	roleColumn      = "role"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

var (
	errQueryBuild   = errors.New("failed to build query")
	errQueryExecute = errors.New("failed to execute query")
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
