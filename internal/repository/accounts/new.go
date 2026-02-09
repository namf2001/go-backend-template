package accounts

import (
	"context"
	"database/sql"

	"github.com/namf2001/go-backend-template/internal/model"
)

type Repository interface {
	// Create creates a new oauth account
	Create(ctx context.Context, account model.Account) (model.Account, error)

	// GetByProvider retrieves an account by provider and provider account id
	GetByProvider(ctx context.Context, provider, providerAccountID string) (model.Account, error)

	// GetByUserID retrieves all accounts for a user
	GetByUserID(ctx context.Context, userID int64) ([]model.Account, error)

	// Delete deletes an account by provider and provider account id
	Delete(ctx context.Context, provider, providerAccountID string) error
}

type impl struct {
	db *sql.DB
}

func New(db *sql.DB) Repository {
	return impl{
		db: db,
	}
}
