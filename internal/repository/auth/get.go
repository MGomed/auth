package auth

import (
	"context"
	"fmt"

	service_model "github.com/MGomed/auth/internal/model"
	repo_model "github.com/MGomed/auth/internal/repository/model"

	sq "github.com/Masterminds/squirrel"
)

// GetUser gets a user in Postgres DB by id
func (a *repository) GetUser(ctx context.Context, id int64) (*service_model.User, error) {
	builder := sq.Select(idColumn, nameColumn, emailColumn, roleColumn, createdAtColumn, updatedAtColumn).
		From(authTable).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("%w - %v : %w", errQueryBuild, query, err)
	}

	var user = &repo_model.User{}
	err = a.pool.QueryRow(ctx, query, args...).
		Scan(user)
	if err != nil {
		return nil, fmt.Errorf("%w - %v : %w", errQueryExecute, query, err)
	}

	return repo_model.ToUserFromRepo(user), nil
}
