package auth

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	consts "github.com/MGomed/auth/consts"
	errors "github.com/MGomed/auth/internal/repository/errors"
	db "github.com/MGomed/auth/pkg/client/db"
)

// DeleteUser deletes a user in Postgres DB by id
func (a *repository) DeleteUser(ctx context.Context, id int64) (int64, error) {
	builder := sq.Delete(consts.AuthTable).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{consts.IDColumn: id})

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, fmt.Errorf("%w - %v : %w", errors.ErrQueryBuild, query, err)
	}

	q := db.Query{
		Name:     "user_repo.Delete",
		QueryRaw: query,
	}

	res, err := a.dbc.DB().Exec(ctx, q, args...)
	if err != nil {
		return 0, fmt.Errorf("%w - %v : %w", errors.ErrQueryExecute, query, err)
	}

	return res.RowsAffected(), nil
}
