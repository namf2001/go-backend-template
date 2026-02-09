package auth

import (
	"context"

	"github.com/namf2001/go-backend-template/internal/repository"
)

type Controller interface {
	// Login handles manual login
	Login(ctx context.Context, input ValidationInput) (string, error)

	// Register handles manual registration
	Register(ctx context.Context, input RegisterInput) (string, error)

	// OAuthLogin handles oauth login/registration
	OAuthLogin(ctx context.Context, input OAuthInput) (string, error)
}

type impl struct {
	repo repository.Registry
}

func New(repo repository.Registry) Controller {
	return impl{
		repo: repo,
	}
}
