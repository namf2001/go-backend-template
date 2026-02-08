package accounts

import (
	"context"
	"database/sql"

	"github.com/namf2001/go-backend-template/internal/model"
)

type Repository interface {
	// Create creates a new account
	Create(ctx context.Context, account model.Account) (model.Account, error)

	// GetByID retrieves an account by ID
	GetByID(ctx context.Context, id int64) (model.Account, error)

	// GetByUserName retrieves an account by username
	GetByUserName(ctx context.Context, username string) (model.Account, error)

	// List retrieves accounts with optional filters
	List(ctx context.Context, filters ListFilters) ([]model.Account, error)

	// Update updates an existing account
	Update(ctx context.Context, account model.Account) error

	// Delete deletes an account by ID
	Delete(ctx context.Context, id int64) error

	// CountAccount returns the total number of accounts
	CountAccount(ctx context.Context) (int64, error)
}
type impl struct {
	db *sql.DB
}

func New(db *sql.DB) Repository {
	return impl{
		db: db,
	}
}
