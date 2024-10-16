package postgres

import (
	"context"
	"errors"
	"fmt"
	"time"

	domain "github.com/MGomed/auth/internal/domain"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgtype"
	pgxpool "github.com/jackc/pgx/v4/pgxpool"
)

const (
	table = "auth"
)

var (
	errQueryBuild   = errors.New("failed to build query")
	errQueryExecute = errors.New("failed to execute query")
)

type PostgresConfig interface {
	DSN() string
}

type adapter struct {
	ctx  context.Context
	pool *pgxpool.Pool
}

func NewAdapter(ctx context.Context, config PostgresConfig) (*adapter, error) {
	pool, err := pgxpool.Connect(ctx, config.DSN())
	if err != nil {
		return nil, err
	}

	return &adapter{
		ctx:  ctx,
		pool: pool,
	}, nil
}

func (a *adapter) CreateUser(ctx context.Context, info *domain.UserInfo) (int, error) {
	builder := sq.Insert(table).
		PlaceholderFormat(sq.Dollar).
		Columns("name", "email", "password", "role").
		Values(info.Name, info.Email, info.Password, info.Role).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, fmt.Errorf("%w - %v : %w", errQueryBuild, query, err)
	}

	var userID int
	err = a.pool.QueryRow(a.ctx, query, args...).Scan(&userID)
	if err != nil {
		return 0, fmt.Errorf("%w - %v : %w", errQueryExecute, query, err)
	}

	return userID, nil
}

func (a *adapter) GetUser(ctx context.Context, id int) (*domain.UserInfo, error) {
	builder := sq.Select("name", "email", "role", "created_at", "updated_at").
		From(table).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": id})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("%w - %v : %w", errQueryBuild, query, err)
	}

	var info = &domain.UserInfo{}
	var updatedTime pgtype.Timestamp
	err = a.pool.QueryRow(ctx, query, args...).
		Scan(&info.Name, &info.Email, &info.Role, &info.CreatedAt, &updatedTime)
	if err != nil {
		return nil, fmt.Errorf("%w - %v : %w", errQueryExecute, query, err)
	}

	info.UpdatedAt = updatedTime.Time

	return info, nil
}

func (a *adapter) UpdateUser(ctx context.Context, info *domain.UserInfo) (int, error) {
	builder := sq.Update(table).
		PlaceholderFormat(sq.Dollar).
		Set("name", info.Name).
		Set("password", info.Password).
		Set("role", info.Role).
		Set("updated_at", time.Now()).
		Where(sq.Eq{"id": info.ID})

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, fmt.Errorf("%w - %v : %w", errQueryBuild, query, err)
	}

	res, err := a.pool.Exec(a.ctx, query, args...)
	if err != nil {
		return 0, fmt.Errorf("%w - %v : %w", errQueryExecute, query, err)
	}

	return int(res.RowsAffected()), nil
}

func (a *adapter) DeleteUser(ctx context.Context, id int) (int, error) {
	builder := sq.Delete(table).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": id})

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, fmt.Errorf("%w - %v : %w", errQueryBuild, query, err)
	}

	res, err := a.pool.Exec(a.ctx, query, args...)
	if err != nil {
		return 0, fmt.Errorf("%w - %v : %w", errQueryExecute, query, err)
	}

	return int(res.RowsAffected()), nil
}
