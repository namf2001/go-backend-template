# Phase 3: Domain Models & Repository Layer

## Mục Tiêu
- Tạo domain models
- Implement repository interface
- Tạo database queries với raw SQL

---

## Bước 1: Tạo Domain Model

### File: `internal/model/user.go`

```go
cat > internal/model/user.go << 'EOF'
package model

import "time"

// User represents a user in the system
type User struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Validate validates user data
func (u *User) Validate() error {
	if u.Email == "" {
		return ErrInvalidEmail
	}
	if u.Name == "" {
		return ErrInvalidName
	}
	return nil
}
EOF
```

### File: `internal/model/errors.go`

```go
cat > internal/model/errors.go << 'EOF'
package model

import "errors"

var (
	ErrInvalidEmail = errors.New("invalid email")
	ErrInvalidName  = errors.New("invalid name")
)
EOF
```

---

## Bước 2: Implement Repository Interface

### File: `internal/repository/users/repository.go`

```go
cat > internal/repository/users/repository.go << 'EOF'
package users

import (
	"context"

	"github.com/yourusername/go-backend-template/internal/model"
)

// Repository defines the interface for user data access
type Repository interface {
	// Create creates a new user
	Create(ctx context.Context, user model.User) (model.User, error)
	
	// GetByID retrieves a user by ID
	GetByID(ctx context.Context, id int64) (model.User, error)
	
	// GetByEmail retrieves a user by email
	GetByEmail(ctx context.Context, email string) (model.User, error)
	
	// List retrieves users with optional filters
	List(ctx context.Context, filters ListFilters) ([]model.User, error)
	
	// Update updates an existing user
	Update(ctx context.Context, user model.User) error
	
	// Delete deletes a user by ID
	Delete(ctx context.Context, id int64) error
}

// ListFilters represents filters for listing users
type ListFilters struct {
	Limit  int
	Offset int
	Email  string
}
EOF
```

---

## Bước 3: Implement Repository

### File: `internal/repository/users/postgres.go`

```go
cat > internal/repository/users/postgres.go << 'EOF'
package users

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"github.com/yourusername/go-backend-template/internal/model"
	apperrors "github.com/yourusername/go-backend-template/internal/pkg/errors"
)

type postgresRepo struct {
	db *sql.DB
}

// NewPostgresRepository creates a new PostgreSQL repository
func NewPostgresRepository(db *sql.DB) Repository {
	return &postgresRepo{db: db}
}

// Create creates a new user
func (r *postgresRepo) Create(ctx context.Context, user model.User) (model.User, error) {
	query := `
		INSERT INTO users (email, name)
		VALUES ($1, $2)
		RETURNING id, email, name, created_at, updated_at
	`

	var created model.User
	err := r.db.QueryRowContext(ctx, query, user.Email, user.Name).Scan(
		&created.ID,
		&created.Email,
		&created.Name,
		&created.CreatedAt,
		&created.UpdatedAt,
	)

	if err != nil {
		// Check for unique constraint violation
		if isUniqueViolation(err) {
			return model.User{}, apperrors.AlreadyExists("user with this email already exists")
		}
		return model.User{}, errors.Wrap(err, "failed to create user")
	}

	return created, nil
}

// GetByID retrieves a user by ID
func (r *postgresRepo) GetByID(ctx context.Context, id int64) (model.User, error) {
	query := `
		SELECT id, email, name, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	var user model.User
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Email,
		&user.Name,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return model.User{}, apperrors.NotFound("user not found")
	}
	if err != nil {
		return model.User{}, errors.Wrap(err, "failed to get user")
	}

	return user, nil
}

// GetByEmail retrieves a user by email
func (r *postgresRepo) GetByEmail(ctx context.Context, email string) (model.User, error) {
	query := `
		SELECT id, email, name, created_at, updated_at
		FROM users
		WHERE email = $1
	`

	var user model.User
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Email,
		&user.Name,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return model.User{}, apperrors.NotFound("user not found")
	}
	if err != nil {
		return model.User{}, errors.Wrap(err, "failed to get user by email")
	}

	return user, nil
}

// List retrieves users with filters
func (r *postgresRepo) List(ctx context.Context, filters ListFilters) ([]model.User, error) {
	query := `
		SELECT id, email, name, created_at, updated_at
		FROM users
		WHERE 1=1
	`
	args := []interface{}{}
	argPos := 1

	// Add email filter if provided
	if filters.Email != "" {
		query += ` AND email ILIKE $` + string(rune(argPos))
		args = append(args, "%"+filters.Email+"%")
		argPos++
	}

	// Add ordering
	query += ` ORDER BY created_at DESC`

	// Add pagination
	if filters.Limit > 0 {
		query += ` LIMIT $` + string(rune(argPos))
		args = append(args, filters.Limit)
		argPos++
	}
	if filters.Offset > 0 {
		query += ` OFFSET $` + string(rune(argPos))
		args = append(args, filters.Offset)
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to list users")
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.Name,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, errors.Wrap(err, "failed to scan user")
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "error iterating users")
	}

	return users, nil
}

// Update updates an existing user
func (r *postgresRepo) Update(ctx context.Context, user model.User) error {
	query := `
		UPDATE users
		SET email = $1, name = $2
		WHERE id = $3
	`

	result, err := r.db.ExecContext(ctx, query, user.Email, user.Name, user.ID)
	if err != nil {
		if isUniqueViolation(err) {
			return apperrors.AlreadyExists("user with this email already exists")
		}
		return errors.Wrap(err, "failed to update user")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "failed to get rows affected")
	}

	if rowsAffected == 0 {
		return apperrors.NotFound("user not found")
	}

	return nil
}

// Delete deletes a user by ID
func (r *postgresRepo) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM users WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return errors.Wrap(err, "failed to delete user")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "failed to get rows affected")
	}

	if rowsAffected == 0 {
		return apperrors.NotFound("user not found")
	}

	return nil
}

// isUniqueViolation checks if error is a unique constraint violation
func isUniqueViolation(err error) bool {
	// PostgreSQL unique violation error code is 23505
	return err != nil && (err.Error() == "pq: duplicate key value violates unique constraint \"users_email_key\"" ||
		err.Error() == "UNIQUE constraint failed")
}
EOF
```

---

## Bước 4: Implement Validator

### File: `internal/pkg/validator/validator.go`

```go
cat > internal/pkg/validator/validator.go << 'EOF'
package validator

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// Validate validates a struct
func Validate(data interface{}) error {
	return validate.Struct(data)
}

// ValidationErrors converts validator errors to readable format
func ValidationErrors(err error) map[string]string {
	errors := make(map[string]string)
	
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			errors[e.Field()] = getErrorMessage(e)
		}
	}
	
	return errors
}

func getErrorMessage(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	case "min":
		return "Value is too short"
	case "max":
		return "Value is too long"
	default:
		return "Invalid value"
	}
}
EOF
```

---

## Kiểm Tra

Tạo file test đơn giản:

### File: `internal/repository/users/postgres_test.go`

```go
cat > internal/repository/users/postgres_test.go << 'EOF'
package users

import (
	"testing"
)

func TestRepository(t *testing.T) {
	// TODO: Add integration tests
	t.Skip("Integration tests require database setup")
}
EOF
```

Chạy test:

```bash
go test ./internal/repository/users/...
```

---

## Kết Quả Mong Đợi

✅ Domain models đã được tạo  
✅ Repository interface đã được định nghĩa  
✅ PostgreSQL repository implementation hoàn chỉnh  
✅ Validator utilities đã sẵn sàng  

---

## Tiếp Theo

Chuyển sang **Phase 4: Controller & Handler** để:
- Implement business logic layer (Controller)
- Tạo REST API handlers
- Setup routing
