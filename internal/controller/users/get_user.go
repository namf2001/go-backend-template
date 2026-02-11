package users

import (
	"context"

	"github.com/namf2001/go-backend-template/internal/model"
)

// GetUser show a user by ID.
func (i impl) GetUser(ctx context.Context, id int64) (model.User, error) {
	return i.repo.User().GetByID(ctx, id)
}
