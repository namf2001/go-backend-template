package users

import "context"

// DeleteUser deletes a user by ID.
func (i impl) DeleteUser(ctx context.Context, id int64) error {
	return i.repo.User().Delete(ctx, id)
}
