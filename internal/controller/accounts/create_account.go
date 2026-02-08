package accounts

import (
	"context"

	"github.com/namf2001/go-backend-template/internal/model"
)

type CreateAccountInput struct {
	UserID   int64
	Username string
	Password string
}

// CreateAccount retrieves an account by ID
func (i impl) CreateAccount(ctx context.Context, input CreateAccountInput) (model.Account, error) {
	//TODO implement me
	panic("implement me")
}
