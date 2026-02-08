package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

// ServerConfig represents the server configuration
type ServerConfig struct {
	Port string
	Env  string
}

// DatabaseConfig represents the database configuration
type DatabaseConfig struct {
	Host         string `envconfig:"DB_HOST" default:"localhost"`
	Port         string `envconfig:"DB_PORT" default:"5432"`
	User         string `envconfig:"DB_USER" default:"postgres"`
	Password     string `envconfig:"DB_PASSWORD" default:"postgres"`
	DBName       string `envconfig:"DB_NAME" default:"go_backend_db"`
	SSLMode      string `envconfig:"DB_SSL_MODE" default:"disable"`
	MaxOpenConns int
	MaxIdleConns int
}

// Config represents the application configuration
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

// Load loads the configuration from environment variables
func Load() (*Config, error) {
	// Load .env file if exists
	_ = godotenv.Load()

	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to process config")
	}
	return &cfg, nil
}
