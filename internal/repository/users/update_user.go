package users

import (
	"context"

	"github.com/namf2001/go-backend-template/internal/model"
	pkgerrors "github.com/pkg/errors"
)

var (
	ErrDuplicateEmail = pkgerrors.New("user with this email already exists")
)

// Update implements Repository.
func (i impl) Update(ctx context.Context, user model.User) error {
	query := `
		UPDATE users
		SET email = $1, name = $2, password = $3, image = $4, "emailVerified" = $5
		WHERE id = $6
	`

	result, err := i.db.ExecContext(ctx, query, user.Email, user.Name, user.Password, user.Image, user.EmailVerified, user.ID)
	if err != nil {
		if err.Error() == "pq: duplicate key value violates unique constraint \"users_email_key\"" ||
			err.Error() == "UNIQUE constraint failed" {
			return pkgerrors.WithStack(ErrDuplicateEmail)
		}
		return pkgerrors.WithStack(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return pkgerrors.WithStack(err)
	}

	if rowsAffected == 0 {
		return pkgerrors.WithStack(ErrNotFound)
	}

	return nil
}
