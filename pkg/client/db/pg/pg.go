package pg

import (
	"context"
	"fmt"
	"log"

	pgxscan "github.com/georgysavva/scany/pgxscan"
	pgconn "github.com/jackc/pgconn"
	pgx "github.com/jackc/pgx/v4"
	pgxpool "github.com/jackc/pgx/v4/pgxpool"

	db "github.com/MGomed/auth/pkg/client/db"
	prettier "github.com/MGomed/auth/pkg/client/db/prettier"
)

type key string

const (
	// TxKey is transaction key in context
	TxKey key = "tx"
)

type pg struct {
	log *log.Logger
	dbc *pgxpool.Pool
}

// NewDB id pg struct constructor
func NewDB(dbc *pgxpool.Pool) db.DB {
	return &pg{
		dbc: dbc,
	}
}

// ScanOne is wrapper for scanning one row into struct from db
func (p *pg) ScanOne(ctx context.Context, dest interface{}, q db.Query, args ...interface{}) error {
	p.logQuery(q, args...)

	row, err := p.Query(ctx, q, args...)
	if err != nil {
		return err
	}

	return pgxscan.ScanOne(dest, row)
}

// ScanAll is wrapper for scanning more then one row into struct from db
func (p *pg) ScanAll(ctx context.Context, dest interface{}, q db.Query, args ...interface{}) error {
	p.logQuery(q, args...)

	rows, err := p.Query(ctx, q, args...)
	if err != nil {
		return err
	}

	return pgxscan.ScanAll(dest, rows)
}

// Exec is wrapper for db Exec() func
func (p *pg) Exec(ctx context.Context, q db.Query, args ...interface{}) (pgconn.CommandTag, error) {
	p.logQuery(q, args...)

	tx, ok := ctx.Value(TxKey).(pgx.Tx)
	if ok {
		return tx.Exec(ctx, q.QueryRaw, args...)
	}

	return p.dbc.Exec(ctx, q.QueryRaw, args...)
}

// Query is wrapper for db Query() func
func (p *pg) Query(ctx context.Context, q db.Query, args ...interface{}) (pgx.Rows, error) {
	p.logQuery(q, args...)

	tx, ok := ctx.Value(TxKey).(pgx.Tx)
	if ok {
		return tx.Query(ctx, q.QueryRaw, args...)
	}

	return p.dbc.Query(ctx, q.QueryRaw, args...)
}

// QueryRow is wrapper for db QueryRow() func
func (p *pg) QueryRow(ctx context.Context, q db.Query, args ...interface{}) pgx.Row {
	p.logQuery(q, args...)

	tx, ok := ctx.Value(TxKey).(pgx.Tx)
	if ok {
		return tx.QueryRow(ctx, q.QueryRaw, args...)
	}

	return p.dbc.QueryRow(ctx, q.QueryRaw, args...)
}

// BeginTx create new transaction object
func (p *pg) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error) {
	return p.dbc.BeginTx(ctx, txOptions)
}

// Ping checks db connection
func (p *pg) Ping(ctx context.Context) error {
	return p.dbc.Ping(ctx)
}

// Close closes db connection
func (p *pg) Close() {
	p.dbc.Close()
}

// MakeContextTx wraps transaction into context
func MakeContextTx(ctx context.Context, tx pgx.Tx) context.Context {
	return context.WithValue(ctx, TxKey, tx)
}

func (p *pg) logQuery(q db.Query, args ...interface{}) {
	prettyQuery := prettier.Pretty(q.QueryRaw, prettier.PlaceholderDollar, args...)
	p.log.Println(
		fmt.Sprintf("sql: %s", q.Name),
		fmt.Sprintf("query: %s", prettyQuery),
	)
}
