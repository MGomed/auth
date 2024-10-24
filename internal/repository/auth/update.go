package auth

import (
	"context"
	"fmt"

	service_model "github.com/MGomed/auth/internal/model"

	sq "github.com/Masterminds/squirrel"
)

// UpdateUser updates a user in Postgres DB
func (a *repository) UpdateUser(ctx context.Context, user *service_model.User) (int64, error) {
	builder := sq.Update(authTable).
		PlaceholderFormat(sq.Dollar)

	if user.Name != nil {
		builder = builder.Set(nameColumn, user.Name)
	}

	if user.Password != nil {
		builder = builder.Set(passwordColumn, user.Password)
	}

	if user.Role != nil {
		builder = builder.Set(roleColumn, user.Role)
	}

	if user.UpdatedAt != nil {
		builder = builder.Set(updatedAtColumn, user.UpdatedAt)
	}

	builder = builder.Where(sq.Eq{idColumn: user.ID})

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, fmt.Errorf("%w - %v : %w", errQueryBuild, query, err)
	}

	res, err := a.pool.Exec(ctx, query, args...)
	if err != nil {
		return 0, fmt.Errorf("%w - %v : %w", errQueryExecute, query, err)
	}

	return res.RowsAffected(), nil
}
