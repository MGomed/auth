package auth

import (
	"errors"

	pgxpool "github.com/jackc/pgx/v4/pgxpool"
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
	pool *pgxpool.Pool
}

// NewRepository is repository struct constructor
func NewRepository(pool *pgxpool.Pool) *repository {
	return &repository{
		pool: pool,
	}
}
