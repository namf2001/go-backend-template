package accounts

import (
	"context"

	"github.com/namf2001/go-backend-template/internal/model"
	"github.com/namf2001/go-backend-template/internal/repository"
)

type Controller interface {
	// CreateAccount creates a new account
	CreateAccount(ctx context.Context, input CreateAccountInput) (model.Account, error)
	// GetAccount retrieves an account by ID
	GetAccount(ctx context.Context, id int64) (model.Account, error)
}

// New creates a new users Controller
func New(repo repository.Registry) Controller {
	return impl{
		repo: repo,
	}
}

type impl struct {
	repo repository.Registry
}
