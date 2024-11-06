package db

import (
	"context"

	pgconn "github.com/jackc/pgconn"
	pgx "github.com/jackc/pgx/v4"
)

//go:generate mockgen -destination=./mocks/db_mock.go -package=mocks -source=interfaces.go

// Handler is func which called in transaction
type Handler func(ctx context.Context) error

// Client is wrapper for db conn
type Client interface {
	DB() DB
	Close() error
}

// TxManager called hadler in transaction
type TxManager interface {
	ReadCommitted(ctx context.Context, f Handler) error
}

// Query is wrapper on sql query with it's name
type Query struct {
	Name     string
	QueryRaw string
}

// Transactor is interface for working with transactions
type Transactor interface {
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
}

// SQLExecer compined NamedExecer and QueryExecer
type SQLExecer interface {
	NamedExecer
	QueryExecer
}

// NamedExecer is interface for working with taged db structs
type NamedExecer interface {
	ScanOne(ctx context.Context, dest interface{}, q Query, args ...interface{}) error
	ScanAll(ctx context.Context, dest interface{}, q Query, args ...interface{}) error
}

// QueryExecer is wrapper for casual db quieries
type QueryExecer interface {
	Exec(ctx context.Context, q Query, args ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, q Query, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, q Query, args ...interface{}) pgx.Row
}

// Pinger is checker for db connection
type Pinger interface {
	Ping(ctx context.Context) error
}

// DB is interface for working with db
type DB interface {
	SQLExecer
	Pinger
	Transactor
	Close()
}
