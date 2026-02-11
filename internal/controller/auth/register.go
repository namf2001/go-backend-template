package auth

import (
	"context"

	"github.com/namf2001/go-backend-template/internal/model"
	"github.com/namf2001/go-backend-template/internal/pkg/jwt"
	"github.com/namf2001/go-backend-template/internal/pkg/utils"
	"github.com/namf2001/go-backend-template/internal/repository"
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

	// 2. Create user + account in a single transaction
	user := model.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: hashedPassword,
	}

	var createdUser model.User // declare outside closure to use after tx
	err = i.repo.DoInTx(ctx, func(ctx context.Context, txRepo repository.Registry) error {
		var txErr error
		createdUser, txErr = txRepo.User().Create(ctx, user)
		if txErr != nil {
			return txErr
		}

		_, txErr = txRepo.Account().Create(ctx, model.Account{
			UserID: createdUser.ID,
			Type:   "personal",
		})
		return txErr
	}, nil)
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
