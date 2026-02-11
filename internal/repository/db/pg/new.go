package pg

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
	pkgerrors "github.com/pkg/errors"
)

// NewPool opens a new DB connection pool, pings it and returns a BeginnerExecutor
func NewPool(dsn string, maxOpenConns int, maxIdleConns int) (BeginnerExecutor, error) {
	log.Println("Initializing Postgres connection pool...")

	pool, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, pkgerrors.WithStack(fmt.Errorf("opening DB failed: %w", err))
	}

	pool.SetMaxOpenConns(maxOpenConns)
	pool.SetMaxIdleConns(maxIdleConns)
	pool.SetConnMaxLifetime(29 * time.Minute)

	// Verify connection
	if err := pool.PingContext(context.Background()); err != nil {
		return nil, pkgerrors.WithStack(fmt.Errorf("unable to ping DB: %w", err))
	}

	log.Println("Postgres connection pool initialized successfully")

	return pool, nil
}
