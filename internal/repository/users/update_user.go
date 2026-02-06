package users

import (
	"context"

	"github.com/namf2001/go-backend-template/internal/model"
	apperrors "github.com/namf2001/go-backend-template/internal/pkg/errors"
	"github.com/pkg/errors"
)

// Update implements Repository.
func (i impl) Update(ctx context.Context, user model.User) error {
	query := `
		UPDATE users
		SET email = $1, name = $2
		WHERE id = $3
	`

	result, err := i.db.ExecContext(ctx, query, user.Email, user.Name, user.ID)
	if err != nil {
		if err != nil && (err.Error() == "pq: duplicate key value violates unique constraint \"users_email_key\"" ||
			err.Error() == "UNIQUE constraint failed") {
			return apperrors.AlreadyExists("user with this email already exists")
		}
		return errors.Wrap(err, "failed to update user")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "failed to get rows affected")
	}

	if rowsAffected == 0 {
		return apperrors.NotFound("user not found")
	}

	return nil
}
