package users

import (
	"context"

	"github.com/namf2001/go-backend-template/internal/model"
	"github.com/pkg/errors"
)

// GetUser show a user by ID.
func (i impl) GetUser(ctx context.Context, id int64) (model.User, error) {
	user, err := i.repo.User().GetByID(ctx, id)
	if err != nil {
		return model.User{}, errors.Wrap(err, "failed to get user")
	}

	return user, nil
}
