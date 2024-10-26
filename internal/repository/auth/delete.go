package auth

import (
	"context"
	"fmt"

	"github.com/MGomed/auth/pkg/client/db"
	sq "github.com/Masterminds/squirrel"
)

// DeleteUser deletes a user in Postgres DB by id
func (a *repository) DeleteUser(ctx context.Context, id int64) (int64, error) {
	builder := sq.Delete(authTable).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id})

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, fmt.Errorf("%w - %v : %w", errQueryBuild, query, err)
	}

	q := db.Query{
		Name:     "user_repo.Delete",
		QueryRaw: query,
	}

	res, err := a.dbc.DB().Exec(ctx, q, args...)
	if err != nil {
		return 0, fmt.Errorf("%w - %v : %w", errQueryExecute, query, err)
	}

	return res.RowsAffected(), nil
}
