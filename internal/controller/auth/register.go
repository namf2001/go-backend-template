package auth

import (
	"context"

	"github.com/namf2001/go-backend-template/internal/model"
	"github.com/namf2001/go-backend-template/internal/pkg/jwt"
	"github.com/namf2001/go-backend-template/internal/pkg/utils"
)

type RegisterInput struct {
	Name     string
	Email    string
	Password string
}

// Register performs manual registration
func (i impl) Register(ctx context.Context, input RegisterInput) (string, error) {
	// 1. Hash password
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return "", err
	}

	// 2. Create user
	user := model.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: hashedPassword,
	}

	createdUser, err := i.repo.User().Create(ctx, user)
	if err != nil {
		return "", err
	}

	// 3. Login (generate token)
	token, err := jwt.GenerateToken(createdUser.ID, createdUser.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}
