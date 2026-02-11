package pg

import (
	"context"
	"database/sql"
	"log"
)

// instrumentedDB wraps the *sql.DB to add logging
type instrumentedDB struct {
	*sql.DB
}

// BeginTx begins a transaction with logging
func (i *instrumentedDB) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	log.Println("DB: BeginTx")
	return i.DB.BeginTx(ctx, opts)
}

// ExecContext wraps the base connector with logging
func (i *instrumentedDB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	log.Printf("DB Exec: %s", query)
	return i.DB.ExecContext(ctx, query, args...)
}

// QueryContext wraps the base connector with logging
func (i *instrumentedDB) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	log.Printf("DB Query: %s", query)
	return i.DB.QueryContext(ctx, query, args...)
}

// QueryRowContext wraps the base connector with logging
func (i *instrumentedDB) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	log.Printf("DB QueryRow: %s", query)
	return i.DB.QueryRowContext(ctx, query, args...)
}

// NewInstrumentedDB wraps a *sql.DB with logging and returns a BeginnerExecutor
func NewInstrumentedDB(db *sql.DB) BeginnerExecutor {
	return &instrumentedDB{DB: db}
}
