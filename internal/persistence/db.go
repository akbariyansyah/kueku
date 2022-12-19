package persistence

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"

)

type Sqlx interface {
	UnitOfWork
	Pinger
	sqlx.ExecerContext
	sqlx.QueryerContext
}
type UnitOfWork interface {
	Tx(ctx context.Context, fun func(ctx context.Context, err chan error)) error
}

// Pinger .
type Pinger interface {
	Ping(ctx context.Context) error
}

type DB struct {
	*sqlx.DB
}

func (r *DB) Startup() error               { return nil }
func (r *DB) Shutdown() error              { return r.DB.Close() }
func (r *DB) Ping(_ context.Context) error { return r.DB.Ping() }

func (r *DB) Tx(ctx context.Context, fun func(ctx context.Context, err chan error)) error {
	e := make(chan error, 1)

	if _, ok := ctx.Value("tx").(*sqlx.Tx); !ok {
		tx, err := r.DB.Beginx()
		if err != nil {
			return err
		}
		ctx = context.WithValue(ctx, "tx", tx)

		go fun(ctx, e)

		if err := <-e; err != nil {
			if errTx := tx.Rollback(); errTx != nil {
				err = fmt.Errorf("%q: %w", errTx, err)
			}
			return err
		}
		return tx.Commit()
	}

	go fun(ctx, e)

	return <-e
}

func (r *DB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	tx, ok := ctx.Value("tx").(*sqlx.Tx)
	if !ok {
		return r.DB.ExecContext(ctx, query, args...)
	}
	return tx.ExecContext(ctx, query, args...)
}

func (r *DB) QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row {
	tx, ok := ctx.Value("tx").(*sqlx.Tx)
	if !ok {
		return r.DB.QueryRowxContext(ctx, query, args...)
	}
	return tx.QueryRowxContext(ctx, query, args...)
}
