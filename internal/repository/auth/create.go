package auth

import (
	"context"
	"fmt"

	service_model "github.com/MGomed/auth/internal/model"
	"github.com/MGomed/auth/pkg/client/db"
	sq "github.com/Masterminds/squirrel"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser creates a user in Postgres DB
func (a *repository) CreateUser(ctx context.Context, user *service_model.UserCreate) (int64, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(user.Password, bcrypt.MinCost)
	if err != nil {
		return 0, fmt.Errorf("failed to hash password: %w", err)
	}

	builder := sq.Insert(authTable).
		PlaceholderFormat(sq.Dollar).
		Columns(nameColumn, emailColumn, passwordColumn, roleColumn).
		Values(user.Name, user.Email, string(hashedPassword), user.Role).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, fmt.Errorf("%w - %v : %w", errQueryBuild, query, err)
	}

	q := db.Query{
		Name:     "user_repo.Create",
		QueryRaw: query,
	}

	var userID int64
	err = a.dbc.DB().QueryRow(ctx, q, args...).Scan(&userID)
	if err != nil {
		return 0, fmt.Errorf("%w - %v : %w", errQueryExecute, query, err)
	}

	return userID, nil
}
