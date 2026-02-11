package auth

import (
	"context"

	"github.com/namf2001/go-backend-template/internal/pkg/jwt"
	"github.com/namf2001/go-backend-template/internal/pkg/utils"
)

type ValidationInput struct {
	Email    string
	Password string
}

// Login performs manual login
func (i impl) Login(ctx context.Context, input ValidationInput) (string, error) {
	// 1. Get user by email
	user, err := i.repo.User().GetByEmail(ctx, input.Email)
	if err != nil {
		return "", err
	}

	// 2. Validate password
	if err := utils.VerifyPassword(user.Password, input.Password); err != nil {
		return "", err
	}

	// 3. Generate Token
	token, err := jwt.GenerateToken(user.ID, user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}
