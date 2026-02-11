package users

import (
	"context"

	"github.com/namf2001/go-backend-template/internal/model"
	"github.com/namf2001/go-backend-template/internal/repository/users"
	pkgerrors "github.com/pkg/errors"
)

// ListFilters represents input for listing users
type ListFilters struct {
	Limit  int
	Offset int
	Email  string
}

// ListUsers lists users based on the provided filters
func (i impl) ListUsers(ctx context.Context, filters ListFilters) ([]model.User, int64, error) {
	repoFilters := users.ListFilters{
		Limit:  filters.Limit,
		Offset: filters.Offset,
		Email:  filters.Email,
	}

	userList, err := i.repo.User().List(ctx, repoFilters)
	if err != nil {
		return []model.User{}, 0, pkgerrors.WithStack(err)
	}

	totalUser, err := i.repo.User().CountUser(ctx)
	if err != nil {
		return []model.User{}, 0, pkgerrors.WithStack(err)
	}

	return userList, totalUser, nil
}
