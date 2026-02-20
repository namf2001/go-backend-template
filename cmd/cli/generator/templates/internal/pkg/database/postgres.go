package database

import (
	"context"
	"fmt"

	"github.com/namf2001/go-backend-template/config"
	"github.com/namf2001/go-backend-template/internal/repository/db/pg"
)

// NewPostgresConnection creates a new PostgreSQL connection and returns a BeginnerExecutor
func NewPostgresConnection() (pg.BeginnerExecutor, error) {
	cfg := config.GetConfig()
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.GetString("DB_HOST"), cfg.GetString("DB_PORT"), cfg.GetString("DB_USER"), cfg.GetString("DB_PASSWORD"), cfg.GetString("DB_NAME"), cfg.GetString("DB_SSL_MODE"),
	)

	return pg.NewPool(
		dsn,
		cfg.GetInt("DB_MAX_OPEN_CONNS"),
		cfg.GetInt("DB_MAX_IDLE_CONNS"),
	)
}

// CheckConnection verifies database connectivity
func CheckConnection(db pg.BeginnerExecutor) error {
	if err := db.PingContext(context.Background()); err != nil {
		return fmt.Errorf("database connection check failed: %w", err)
	}
	return nil
}
