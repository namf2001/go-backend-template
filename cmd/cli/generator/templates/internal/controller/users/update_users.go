package users

import (
	"context"

	"github.com/namf2001/go-backend-template/internal/pkg/validator"
	pkgerrors "github.com/pkg/errors"
)

// UpdateUserInput represents input for updating a user
type UpdateUserInput struct {
	Email string `validate:"omitempty,email"`
	Name  string `validate:"omitempty,min=2,max=100"`
}

// UpdateUser implements Controller.
func (i impl) UpdateUser(ctx context.Context, id int64, input UpdateUserInput) error {
	// Validate input
	if err := validator.Validate(input); err != nil {
		return pkgerrors.WithStack(err)	
	}

	// Get existing user
	user, err := i.repo.User().GetByID(ctx, id)
	if err != nil {
		return pkgerrors.WithStack(err)
	}

	// Update fields
	if input.Email != "" {
		user.Email = input.Email
	}
	if input.Name != "" {
		user.Name = input.Name
	}

	// Save changes
	if err := i.repo.User().Update(ctx, user); err != nil {
		return pkgerrors.WithStack(err)
	}

	return nil
}
