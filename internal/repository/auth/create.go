package auth

import (
	"context"
	"fmt"

	service_model "github.com/MGomed/auth/internal/model"
	sq "github.com/Masterminds/squirrel"
)

// CreateUser creates a user in Postgres DB
func (a *repository) CreateUser(ctx context.Context, user *service_model.User) (int64, error) {
	builder := sq.Insert(authTable).
		PlaceholderFormat(sq.Dollar).
		Columns(nameColumn, emailColumn, passwordColumn, roleColumn).
		Values(user.Name, user.Email, user.Password, user.Role).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, fmt.Errorf("%w - %v : %w", errQueryBuild, query, err)
	}

	var userID int64
	err = a.pool.QueryRow(ctx, query, args...).Scan(&userID)
	if err != nil {
		return 0, fmt.Errorf("%w - %v : %w", errQueryExecute, query, err)
	}

	return userID, nil
}
