package users

import (
	"context"

	"github.com/namf2001/go-backend-template/internal/model"
	apperrors "github.com/namf2001/go-backend-template/internal/pkg/errors"
	"github.com/namf2001/go-backend-template/internal/pkg/validator"
	"github.com/pkg/errors"
)

// CreateUserInput represents input for creating a user
type CreateUserInput struct {
	Email string `validate:"required,email"`
	Name  string `validate:"required,min=2,max=100"`
}

// CreateUser creates a new user
func (i impl) CreateUser(ctx context.Context, input CreateUserInput) (model.User, error) {
	var UserOutput model.User
	// Validate input
	if err := validator.Validate(input); err != nil {
		return UserOutput, apperrors.InvalidInput("validation failed")
	}

	// Check if a user already exists
	_, err := i.repo.User().GetByEmail(ctx, input.Email)
	if err == nil {
		return UserOutput, apperrors.AlreadyExists("user with this email already exists")
	}

	// Create user
	user := model.User{
		Email: input.Email,
		Name:  input.Name,
	}

	created, err := i.repo.User().Create(ctx, user)
	if err != nil {
		return UserOutput, errors.Wrap(err, "failed to create user")
	}

	return created, nil
}
