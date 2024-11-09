package auth

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	consts "github.com/MGomed/auth/consts"
	service_model "github.com/MGomed/auth/internal/model"
	repo_errors "github.com/MGomed/auth/internal/storage/repository/errors"
	db "github.com/MGomed/common/pkg/client/db"
)

// UpdateUser updates a user in Postgres DB
func (a *repository) UpdateUser(ctx context.Context, user *service_model.UserUpdate) (int64, error) {
	builder := sq.Update(consts.AuthTable)

	if user.Name != nil {
		builder = builder.Set(consts.NameColumn, user.Name)
	}

	if user.Role != nil {
		builder = builder.Set(consts.RoleColumn, user.Role)
	}

	builder = builder.Where(sq.Eq{consts.IDColumn: user.ID}).PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, fmt.Errorf("%w - %v : %w", repo_errors.ErrQueryBuild, query, err)
	}

	q := db.Query{
		Name:     "user_repo.Update",
		QueryRaw: query,
	}

	res, err := a.dbc.DB().Exec(ctx, q, args...)
	if err != nil {
		return 0, fmt.Errorf("%w - %v : %w", repo_errors.ErrQueryExecute, query, err)
	}

	return res.RowsAffected(), nil
}
