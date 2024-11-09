package auth

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	bcrypt "golang.org/x/crypto/bcrypt"

	consts "github.com/MGomed/auth/consts"
	service_model "github.com/MGomed/auth/internal/model"
	repo_errors "github.com/MGomed/auth/internal/storage/repository/errors"
	db "github.com/MGomed/common/pkg/client/db"
)

// CreateUser creates a user in Postgres DB
func (a *repository) CreateUser(ctx context.Context, user *service_model.UserCreate) (int64, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(user.Password, bcrypt.MinCost)
	if err != nil {
		return 0, fmt.Errorf("failed to hash password: %w", err)
	}

	builder := sq.Insert(consts.AuthTable).
		PlaceholderFormat(sq.Dollar).
		Columns(consts.NameColumn, consts.EmailColumn, consts.PasswordColumn, consts.RoleColumn).
		Values(user.Name, user.Email, string(hashedPassword), user.Role).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, fmt.Errorf("%w - %v : %w", repo_errors.ErrQueryBuild, query, err)
	}

	q := db.Query{
		Name:     "user_repo.Create",
		QueryRaw: query,
	}

	var userID int64
	err = a.dbc.DB().QueryRow(ctx, q, args...).Scan(&userID)
	if err != nil {
		return 0, fmt.Errorf("%w - %v : %w", repo_errors.ErrQueryExecute, query, err)
	}

	return userID, nil
}
