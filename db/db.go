package db

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// Handler is a database handler.
type Handler interface {
	Rebind(string) string

	Select(interface{}, string, ...interface{}) error
	Get(interface{}, string, ...interface{}) error
	Queryx(string, ...interface{}) (*sqlx.Rows, error)
	QueryRowx(string, ...interface{}) *sqlx.Row
	Exec(string, ...interface{}) (sql.Result, error)

	SelectContext(context.Context, interface{}, string, ...interface{}) error
	GetContext(context.Context, interface{}, string, ...interface{}) error
	QueryxContext(context.Context, string, ...interface{}) (*sqlx.Rows, error)
	QueryRowxContext(context.Context, string, ...interface{}) *sqlx.Row
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
}

// New creates a new set of queries with the given database handler.
func New(db *sqlx.DB) *Queries {
	return &Queries{db: db}
}

// Queries is a set of queries.
type Queries struct {
	db Handler
}

// WithTx returns a new set of queries with the given transaction.
func (q *Queries) WithTx(tx *sqlx.Tx) *Queries {
	return &Queries{
		db: tx,
	}
}
