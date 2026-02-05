# Phase 2: Database & Infrastructure

## Mục Tiêu
- Tạo database migrations
- Setup database connection pool
- Implement error handling system
- Tạo response utilities

---

## Bước 1: Tạo Database Migration

### File: `migrations/001_create_users_table.sql`

```bash
cat > migrations/001_create_users_table.sql << 'EOF'
-- Create users table
CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Create index on email for faster lookups
CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);

-- Create updated_at trigger function
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Create trigger to auto-update updated_at
CREATE TRIGGER update_users_updated_at 
    BEFORE UPDATE ON users 
    FOR EACH ROW 
    EXECUTE FUNCTION update_updated_at_column();
EOF
```

---

## Bước 2: Implement Database Connection

### File: `internal/pkg/database/postgres.go`

```go
cat > internal/pkg/database/postgres.go << 'EOF'
package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type Config struct {
	Host         string
	Port         string
	User         string
	Password     string
	DBName       string
	SSLMode      string
	MaxOpenConns int
	MaxIdleConns int
}

// NewPostgresConnection creates a new PostgreSQL connection
func NewPostgresConnection(cfg Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open database connection")
	}

	// Set connection pool settings
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
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
EOF
```

---

## Bước 3: Implement Error Handling

### File: `internal/pkg/errors/errors.go`

```go
cat > internal/pkg/errors/errors.go << 'EOF'
package errors

import (
	"errors"
	"fmt"
)

// Common application errors
var (
	ErrNotFound      = errors.New("resource not found")
	ErrAlreadyExists = errors.New("resource already exists")
	ErrInvalidInput  = errors.New("invalid input")
	ErrInternal      = errors.New("internal server error")
	ErrUnauthorized  = errors.New("unauthorized")
	ErrForbidden     = errors.New("forbidden")
)

// AppError represents an application error with additional context
type AppError struct {
	Code    string
	Message string
	Err     error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

// New creates a new AppError
func New(code, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

// NotFound creates a not found error
func NotFound(message string) *AppError {
	return &AppError{
		Code:    "NOT_FOUND",
		Message: message,
		Err:     ErrNotFound,
	}
}

// AlreadyExists creates an already exists error
func AlreadyExists(message string) *AppError {
	return &AppError{
		Code:    "ALREADY_EXISTS",
		Message: message,
		Err:     ErrAlreadyExists,
	}
}

// InvalidInput creates an invalid input error
func InvalidInput(message string) *AppError {
	return &AppError{
		Code:    "INVALID_INPUT",
		Message: message,
		Err:     ErrInvalidInput,
	}
}

// Internal creates an internal error
func Internal(message string, err error) *AppError {
	return &AppError{
		Code:    "INTERNAL_ERROR",
		Message: message,
		Err:     err,
	}
}
EOF
```

---

## Bước 4: Implement Response Utilities

### File: `internal/pkg/response/response.go`

```go
cat > internal/pkg/response/response.go << 'EOF'
package response

import (
	"encoding/json"
	"net/http"

	apperrors "github.com/yourusername/go-backend-template/internal/pkg/errors"
)

// Response represents a standard API response
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   *ErrorInfo  `json:"error,omitempty"`
}

// ErrorInfo represents error information in response
type ErrorInfo struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// JSON sends a JSON response
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	
	response := Response{
		Success: statusCode < 400,
		Data:    data,
	}
	
	json.NewEncoder(w).Encode(response)
}

// Error sends an error response
func Error(w http.ResponseWriter, err error) {
	var appErr *apperrors.AppError
	var statusCode int
	var errorInfo ErrorInfo

	if errors, ok := err.(*apperrors.AppError); ok {
		appErr = errors
		errorInfo = ErrorInfo{
			Code:    appErr.Code,
			Message: appErr.Message,
		}
		
		// Map error types to HTTP status codes
		switch appErr.Code {
		case "NOT_FOUND":
			statusCode = http.StatusNotFound
		case "ALREADY_EXISTS":
			statusCode = http.StatusConflict
		case "INVALID_INPUT":
			statusCode = http.StatusBadRequest
		case "UNAUTHORIZED":
			statusCode = http.StatusUnauthorized
		case "FORBIDDEN":
			statusCode = http.StatusForbidden
		default:
			statusCode = http.StatusInternalServerError
		}
	} else {
		statusCode = http.StatusInternalServerError
		errorInfo = ErrorInfo{
			Code:    "INTERNAL_ERROR",
			Message: "An unexpected error occurred",
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	
	response := Response{
		Success: false,
		Error:   &errorInfo,
	}
	
	json.NewEncoder(w).Encode(response)
}

// Success sends a success response
func Success(w http.ResponseWriter, data interface{}) {
	JSON(w, http.StatusOK, data)
}

// Created sends a created response
func Created(w http.ResponseWriter, data interface{}) {
	JSON(w, http.StatusCreated, data)
}

// NoContent sends a no content response
func NoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}
EOF
```

---

## Bước 5: Implement Configuration

### File: `config/config.go`

```go
cat > config/config.go << 'EOF'
package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Port string
	Env  string
}

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

// Load loads configuration from environment variables
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

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
EOF
```

---

## Bước 6: Chạy Migration

```bash
# Đảm bảo PostgreSQL đang chạy
make docker-up

# Chờ PostgreSQL khởi động hoàn toàn
sleep 5

# Chạy migration
make migrate-up
```

---

## Kiểm Tra

```bash
# Kiểm tra table đã được tạo
psql -h localhost -U postgres -d go_backend_db -c "\dt"

# Kiểm tra cấu trúc table
psql -h localhost -U postgres -d go_backend_db -c "\d users"
```

---

## Kết Quả Mong Đợi

✅ Database migration đã chạy thành công  
✅ Table `users` đã được tạo  
✅ Database connection package đã sẵn sàng  
✅ Error handling system đã được implement  
✅ Response utilities đã sẵn sàng  

---

## Tiếp Theo

Chuyển sang **Phase 3: Domain Models & Repository** để:
- Tạo domain models
- Implement repository pattern
- Tạo database queries
