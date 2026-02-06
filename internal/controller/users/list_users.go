package users

import (
	"context"

	"github.com/namf2001/go-backend-template/internal/model"
	"github.com/namf2001/go-backend-template/internal/repository/users"
	"github.com/pkg/errors"
)

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
		return []model.User{}, 0, errors.Wrap(err, "failed to list users")
	}

	totalUser, err := i.repo.User().CountUser(ctx)
	if err != nil {
		return []model.User{}, 0, errors.Wrap(err, "failed to count users")
	}

	return userList, totalUser, nil
}
