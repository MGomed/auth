package auth

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	consts "github.com/MGomed/auth/consts"
	service_model "github.com/MGomed/auth/internal/model"
	errors "github.com/MGomed/auth/internal/repository/errors"
	db "github.com/MGomed/auth/pkg/client/db"
)

// UpdateUser updates a user in Postgres DB
func (a *repository) UpdateUser(ctx context.Context, user *service_model.UserUpdate) (int64, error) {
	builder := sq.Update(consts.AuthTable).
		PlaceholderFormat(sq.Dollar)

	if user.Name != nil {
		builder = builder.Set(consts.NameColumn, user.Name)
	}

	if user.Role != nil {
		builder = builder.Set(consts.RoleColumn, user.Role)
	}

	builder = builder.Where(sq.Eq{consts.IDColumn: user.ID})

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, fmt.Errorf("%w - %v : %w", errors.ErrQueryBuild, query, err)
	}

	q := db.Query{
		Name:     "user_repo.Update",
		QueryRaw: query,
	}

	res, err := a.dbc.DB().Exec(ctx, q, args...)
	if err != nil {
		return 0, fmt.Errorf("%w - %v : %w", errors.ErrQueryExecute, query, err)
	}

	return res.RowsAffected(), nil
}
