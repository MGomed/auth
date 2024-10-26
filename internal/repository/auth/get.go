package auth

import (
	"context"
	"fmt"

	service_model "github.com/MGomed/auth/internal/model"
	repo_model "github.com/MGomed/auth/internal/repository/model"
	"github.com/MGomed/auth/pkg/client/db"

	sq "github.com/Masterminds/squirrel"
)

// GetUser gets a user in Postgres DB by id
func (a *repository) GetUser(ctx context.Context, id int64) (*service_model.UserInfo, error) {
	builder := sq.Select(idColumn, nameColumn, emailColumn, roleColumn, passwordColumn, createdAtColumn, updatedAtColumn).
		From(authTable).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("%w - %v : %w", errQueryBuild, query, err)
	}

	q := db.Query{
		Name:     "user_repo.Get",
		QueryRaw: query,
	}

	var user repo_model.UserInfo
	err = a.dbc.DB().ScanOne(ctx, &user, q, args...)
	if err != nil {
		return nil, fmt.Errorf("%w - %v : %w", errQueryExecute, query, err)
	}

	return repo_model.ToUserFromRepo(&user), nil
}
