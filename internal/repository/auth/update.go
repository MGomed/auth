package auth

import (
	"context"
	"fmt"

	service_model "github.com/MGomed/auth/internal/model"
	"github.com/MGomed/auth/pkg/client/db"

	sq "github.com/Masterminds/squirrel"
)

// UpdateUser updates a user in Postgres DB
func (a *repository) UpdateUser(ctx context.Context, user *service_model.UserUpdate) (int64, error) {
	builder := sq.Update(authTable).
		PlaceholderFormat(sq.Dollar)

	if user.Name != nil {
		builder = builder.Set(nameColumn, user.Name)
	}

	if user.Role != nil {
		builder = builder.Set(roleColumn, user.Role)
	}

	builder = builder.Where(sq.Eq{idColumn: user.ID})

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, fmt.Errorf("%w - %v : %w", errQueryBuild, query, err)
	}

	q := db.Query{
		Name:     "user_repo.Update",
		QueryRaw: query,
	}

	res, err := a.dbc.DB().Exec(ctx, q, args...)
	if err != nil {
		return 0, fmt.Errorf("%w - %v : %w", errQueryExecute, query, err)
	}

	return res.RowsAffected(), nil
}
