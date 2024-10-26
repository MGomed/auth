package db

import (
	"context"

	pgconn "github.com/jackc/pgconn"
	pgx "github.com/jackc/pgx/v4"
)

type Handler func(ctx context.Context) error

type Client interface {
	DB() DB
	Close() error
}

type Query struct {
	Name     string
	QueryRaw string
}

type SQLExecer interface {
	NamedExecer
	QueryExecer
}

type NamedExecer interface {
	ScanOne(ctx context.Context, dest interface{}, q Query, args ...interface{}) error
	ScanAll(ctx context.Context, dest interface{}, q Query, args ...interface{}) error
}

type QueryExecer interface {
	Exec(ctx context.Context, q Query, args ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, q Query, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, q Query, args ...interface{}) pgx.Row
}

type Pinger interface {
	Ping(ctx context.Context) error
}

type DB interface {
	SQLExecer
	Pinger
	Close()
}
