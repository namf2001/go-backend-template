package users

import (
	"context"
	"database/sql"

	"github.com/namf2001/go-backend-template/internal/model"
)

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

	// CountUser returns the total number of users
	CountUser(ctx context.Context) (int64, error)
}
type impl struct {
	db *sql.DB
}

func New(db *sql.DB) Repository {
	return impl{
		db: db,
	}
}
