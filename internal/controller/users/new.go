package users

import (
	"context"

	"github.com/namf2001/go-backend-template/internal/model"
	"github.com/namf2001/go-backend-template/internal/repository"
)

// Controller defines the user's controller interface
type Controller interface {
	CreateUser(ctx context.Context, input CreateUserInput) (model.User, error)
	GetUser(ctx context.Context, id int64) (model.User, error)
	ListUsers(ctx context.Context, filters ListFilters) ([]model.User, int64, error)
	UpdateUser(ctx context.Context, id int64, input UpdateUserInput) error
	DeleteUser(ctx context.Context, id int64) error
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
