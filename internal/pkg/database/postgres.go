package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/namf2001/go-backend-template/config"
	"github.com/pkg/errors"
)
// NewPostgresConnection creates a new PostgreSQL connection
func NewPostgresConnection() (*sql.DB, error) {
	cfg := config.GetConfig()
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.GetString("DB_HOST"), cfg.GetString("DB_PORT"), cfg.GetString("DB_USER"), cfg.GetString("DB_PASSWORD"), cfg.GetString("DB_NAME"), cfg.GetString("DB_SSL_MODE"),
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open database connection")
	}

	// Set connection pool settings
	db.SetMaxOpenConns(cfg.GetInt("DB_MAX_OPEN_CONNS"))
	db.SetMaxIdleConns(cfg.GetInt("DB_MAX_IDLE_CONNS"))
	db.SetConnMaxLifetime(time.Hour)

	// Verify connection
	if err := db.Ping(); err != nil {
		return nil, errors.Wrap(err, "failed to ping database")
	}

	return db, nil
}

// CheckConnection verifies database connectivity
func CheckConnection(db *sql.DB) error {
	if err := db.Ping(); err != nil {
		return errors.Wrap(err, "database connection check failed")
	}
	return nil
}
