package users

import (
	"context"

	"github.com/namf2001/go-backend-template/internal/model"
	apperrors "github.com/namf2001/go-backend-template/internal/pkg/errors"
	"github.com/pkg/errors"
)

// Create implements Repository.
func (i impl) Create(ctx context.Context, user model.User) (model.User, error) {
	query := `
		INSERT INTO users (email, name)
		VALUES ($1, $2)
		RETURNING id, email, name, created_at, updated_at
	`

	var created model.User
	err := i.db.QueryRowContext(ctx, query, user.Email, user.Name).Scan(
		&created.ID,
		&created.Email,
		&created.Name,
		&created.CreatedAt,
		&created.UpdatedAt,
	)

	if err != nil {
		// Check for unique constraint violation
		if err.Error() == "pq: duplicate key value violates unique constraint \"users_email_key\"" ||
			err.Error() == "UNIQUE constraint failed" {

			return model.User{}, apperrors.AlreadyExists("user with this email already exists")
		}

		return model.User{}, errors.Wrap(err, "failed to create user")
	}

	return created, nil
}
