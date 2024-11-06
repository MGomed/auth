package transaction

import (
	"context"
	"fmt"

	pgx "github.com/jackc/pgx/v4"
	"github.com/pkg/errors"

	"github.com/MGomed/auth/pkg/client/db"
	"github.com/MGomed/auth/pkg/client/db/pg"
)

type manager struct {
	db db.Transactor
}

// NewTransactionManager is manger struct constructor
func NewTransactionManager(db db.Transactor) db.TxManager {
	return &manager{
		db: db,
	}
}

func (m *manager) transaction(ctx context.Context, opts pgx.TxOptions, fn db.Handler) (err error) {
	tx, ok := ctx.Value(pg.TxKey).(pgx.Tx)
	if ok {
		return fn(ctx)
	}

	tx, err = m.db.BeginTx(ctx, opts)
	if err != nil {
		return fmt.Errorf("can't begin transaction: %w", err)
	}

	ctx = pg.MakeContextTx(ctx, tx)

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic recovered: %v", r)
		}

		if err != nil {
			if errRollback := tx.Rollback(ctx); errRollback != nil {
				err = fmt.Errorf("%w: %w", err, errRollback)
			}

			return
		}

		if nil == err {
			err = tx.Commit(ctx)
			if err != nil {
				err = fmt.Errorf("tx commit failed: %w", err)
			}
		}
	}()

	if err = fn(ctx); err != nil {
		err = errors.Wrap(err, "failed executing code inside transaction")
	}

	return err
}

func (m *manager) ReadCommitted(ctx context.Context, f db.Handler) error {
	txOpts := pgx.TxOptions{IsoLevel: pgx.ReadCommitted}
	return m.transaction(ctx, txOpts, f)
}
