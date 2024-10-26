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

type pg struct {
	log *log.Logger
	dbc *pgxpool.Pool
}

func NewDB(dbc *pgxpool.Pool) db.DB {
	return &pg{
		dbc: dbc,
	}
}

func (p *pg) ScanOne(ctx context.Context, dest interface{}, q db.Query, args ...interface{}) error {
	p.logQuery(q, args...)

	row, err := p.Query(ctx, q, args...)
	if err != nil {
		return err
	}

	return pgxscan.ScanOne(dest, row)
}

func (p *pg) ScanAll(ctx context.Context, dest interface{}, q db.Query, args ...interface{}) error {
	p.logQuery(q, args...)

	rows, err := p.Query(ctx, q, args...)
	if err != nil {
		return err
	}

	return pgxscan.ScanAll(dest, rows)
}

func (p *pg) Exec(ctx context.Context, q db.Query, args ...interface{}) (pgconn.CommandTag, error) {
	p.logQuery(q, args...)

	return p.dbc.Exec(ctx, q.QueryRaw, args...)
}

func (p *pg) Query(ctx context.Context, q db.Query, args ...interface{}) (pgx.Rows, error) {
	p.logQuery(q, args...)

	return p.dbc.Query(ctx, q.QueryRaw, args...)
}

func (p *pg) QueryRow(ctx context.Context, q db.Query, args ...interface{}) pgx.Row {
	p.logQuery(q, args...)

	return p.dbc.QueryRow(ctx, q.QueryRaw, args...)
}

func (p *pg) Ping(ctx context.Context) error {
	return p.dbc.Ping(ctx)
}

func (p *pg) Close() {
	p.dbc.Close()
}

func (p *pg) logQuery(q db.Query, args ...interface{}) {
	prettyQuery := prettier.Pretty(q.QueryRaw, prettier.PlaceholderDollar, args...)
	p.log.Println(
		fmt.Sprintf("sql: %s", q.Name),
		fmt.Sprintf("query: %s", prettyQuery),
	)
}
