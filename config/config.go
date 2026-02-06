package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

// ServerConfig represents the server configuration
type ServerConfig struct {
	Port string
	Env  string
}

// DatabaseConfig represents the database configuration
type DatabaseConfig struct {
	Host         string
	Port         string
	User         string
	Password     string
	DBName       string
	SSLMode      string
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

	maxOpenConns, err := strconv.Atoi(getEnv("DB_MAX_OPEN_CONNS", "25"))
	if err != nil {
		return nil, errors.Wrap(err, "invalid DB_MAX_OPEN_CONNS")
	}

	maxIdleConns, err := strconv.Atoi(getEnv("DB_MAX_IDLE_CONNS", "5"))
	if err != nil {
		return nil, errors.Wrap(err, "invalid DB_MAX_IDLE_CONNS")
	}

	return &Config{
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8080"),
			Env:  getEnv("SERVER_ENV", "development"),
		},
		Database: DatabaseConfig{
			Host:         getEnv("DB_HOST", "localhost"),
			Port:         getEnv("DB_PORT", "5432"),
			User:         getEnv("DB_USER", "postgres"),
			Password:     getEnv("DB_PASSWORD", "postgres"),
			DBName:       getEnv("DB_NAME", "go_backend_db"),
			SSLMode:      getEnv("DB_SSL_MODE", "disable"),
			MaxOpenConns: maxOpenConns,
			MaxIdleConns: maxIdleConns,
		},
	}, nil
}

// getEnv returns the value of an environment variable or a default value if the variable is not set
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
