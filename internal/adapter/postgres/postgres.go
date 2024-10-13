package postgres

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/MGomed/auth/internal/domain"
	sq "github.com/Masterminds/squirrel"
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
	log  *log.Logger
	pool *pgxpool.Pool
}

func NewAdapter(ctx context.Context, log *log.Logger, config PostgresConfig) (*adapter, error) {
	pool, err := pgxpool.Connect(ctx, config.DSN())
	if err != nil {
		return nil, err
	}

	return &adapter{
		ctx:  ctx,
		log:  log,
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
		a.log.Printf("%v - %v : %v", errQueryBuild, query, err)

		return 0, err
	}

	var userID int
	err = a.pool.QueryRow(a.ctx, query, args...).Scan(userID)
	if err != nil {
		a.log.Printf("%v - %v : %v", errQueryExecute, query, err)

		return 0, err
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
		a.log.Printf("%v - %v : %v", errQueryBuild, query, err)

		return nil, err
	}

	var info = &domain.UserInfo{}
	err = a.pool.QueryRow(ctx, query, args...).
		Scan(&info.Name, &info.Email, &info.Role, &info.CreatedAt, &info.UpdatedAt)
	if err != nil {
		a.log.Printf("%v - %v : %v", errQueryExecute, query, err)

		return nil, err
	}

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
		a.log.Printf("%v - %v : %v", errQueryBuild, query, err)

		return 0, err
	}

	res, err := a.pool.Exec(a.ctx, query, args...)
	if err != nil {
		a.log.Printf("%v - %v : %v", errQueryExecute, query, err)

		return 0, err
	}

	return int(res.RowsAffected()), nil
}

func (a *adapter) DeleteUser(ctx context.Context, id int) (int, error) {
	builder := sq.Delete(table).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": id})

	query, args, err := builder.ToSql()
	if err != nil {
		a.log.Printf("%v - %v : %v", errQueryBuild, query, err)

		return 0, err
	}

	res, err := a.pool.Exec(a.ctx, query, args...)
	if err != nil {
		a.log.Printf("%v - %v : %v", errQueryExecute, query, err)

		return 0, err
	}

	return int(res.RowsAffected()), nil
}
