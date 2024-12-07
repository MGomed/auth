package auth

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	consts "github.com/MGomed/auth/consts"
	service_model "github.com/MGomed/auth/internal/model"
	repo_converters "github.com/MGomed/auth/internal/storage/repository/converters"
	repo_errors "github.com/MGomed/auth/internal/storage/repository/errors"
	repo_model "github.com/MGomed/auth/internal/storage/repository/model"
	db "github.com/MGomed/common/client/db"
)

// GetUser gets a user in Postgres DB by id
func (a *repository) GetUser(ctx context.Context, id int64) (*service_model.UserInfo, error) {
	builder := sq.Select(consts.IDColumn, consts.NameColumn, consts.EmailColumn,
		consts.RoleColumn, consts.PasswordColumn, consts.CreatedAtColumn, consts.UpdatedAtColumn).
		From(consts.AuthTable).
		Where(sq.Eq{consts.IDColumn: id}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("%w - %v : %w", repo_errors.ErrQueryBuild, query, err)
	}

	q := db.Query{
		Name:     "user_repo.Get",
		QueryRaw: query,
	}

	var user repo_model.UserInfo
	err = a.dbc.DB().ScanOne(ctx, &user, q, args...)
	if err != nil {
		return nil, fmt.Errorf("%w - %v : %w", repo_errors.ErrQueryExecute, query, err)
	}

	return repo_converters.ToUserFromRepo(&user), nil
}

// GetUserByEmail gets a user in Postgres DB by email
func (a *repository) GetUserByEmail(ctx context.Context, email string) (*service_model.UserInfo, error) {
	builder := sq.Select(consts.IDColumn, consts.NameColumn, consts.EmailColumn,
		consts.RoleColumn, consts.PasswordColumn, consts.CreatedAtColumn, consts.UpdatedAtColumn).
		From(consts.AuthTable).
		Where(sq.Eq{consts.EmailColumn: email}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("%w - %v : %w", repo_errors.ErrQueryBuild, query, err)
	}

	q := db.Query{
		Name:     "user_repo.GetByEmail",
		QueryRaw: query,
	}

	var user repo_model.UserInfo
	err = a.dbc.DB().ScanOne(ctx, &user, q, args...)
	if err != nil {
		return nil, fmt.Errorf("%w - %v : %w", repo_errors.ErrQueryExecute, query, err)
	}

	return repo_converters.ToUserFromRepo(&user), nil
}
