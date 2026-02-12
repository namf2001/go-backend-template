package testdb

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"sync"
	"testing"

	_ "github.com/lib/pq"
	"github.com/namf2001/go-backend-template/internal/repository/db/pg"
	"github.com/stretchr/testify/require"
)

var (
	appDB   *sql.DB
	dbOnce  sync.Once
	initErr error
)

func getDB(t *testing.T) *sql.DB {
	t.Helper()
	dbOnce.Do(func() {
		dsn := os.Getenv("TEST_DB_DSN")
		if dsn == "" {
			// Fallback: build DSN from individual env vars
			host := envOrDefault("DB_HOST", "localhost")
			port := envOrDefault("DB_PORT", "5432")
			user := envOrDefault("DB_USER", "postgres")
			password := envOrDefault("DB_PASSWORD", "postgres")
			dbName := envOrDefault("DB_NAME", "go_backend_test")
			sslMode := envOrDefault("DB_SSL_MODE", "disable")

			dsn = fmt.Sprintf(
				"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
				host, port, user, password, dbName, sslMode,
			)
		}

		var err error
		appDB, err = sql.Open("postgres", dsn)
		if err != nil {
			initErr = fmt.Errorf("failed to open test DB: %w", err)
			return
		}

		appDB.SetMaxOpenConns(5)
		appDB.SetMaxIdleConns(2)

		if err = appDB.PingContext(context.Background()); err != nil {
			initErr = fmt.Errorf("failed to ping test DB: %w", err)
			return
		}
	})

	require.NoError(t, initErr, "test DB initialization failed")
	return appDB
}

// WithTx provides a callback with a *sql.Tx (which satisfies pg.ContextExecutor)
// for running repository tests. The transaction is always rolled back after the
// callback completes, so no data is actually persisted to the database.
// This pattern is adapted from ingenta-be's testent.WithEntTx.
func WithTx(t *testing.T, callback func(tx pg.ContextExecutor)) {
	t.Helper()
	db := getDB(t)

	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	require.NoError(t, err, "failed to begin test transaction")

	defer func() {
		err := tx.Rollback()
		if err != nil && !errors.Is(err, sql.ErrTxDone) {
			t.Errorf("failed to rollback test transaction: %v", err)
		}
	}()

	callback(tx)
}

// LoadTestSQLFile loads a SQL file from the filesystem relative to the caller's
// working directory and executes it within the given transaction.
// This pattern is adapted from ingenta-be's testent.LoadTestSQLFile.
func LoadTestSQLFile(t *testing.T, tx pg.ContextExecutor, filename string) {
	t.Helper()

	body, err := os.ReadFile(filename)
	require.NoError(t, err, "failed to read test SQL file: %s", filename)

	_, err = tx.ExecContext(context.Background(), string(body))
	require.NoError(t, err, "failed to execute test SQL file: %s", filename)
}

func envOrDefault(key, defaultVal string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultVal
}
