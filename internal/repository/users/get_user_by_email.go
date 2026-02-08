package users

import (
	"context"
	"database/sql"

	"github.com/namf2001/go-backend-template/internal/model"
	apperrors "github.com/namf2001/go-backend-template/internal/pkg/errors"
	"github.com/pkg/errors"
)

// GetByEmail implements Repository.
func (i impl) GetByEmail(ctx context.Context, email string) (model.User, error) {
	query := `
		SELECT id, email, name, created_at, updated_at
		FROM users
		WHERE email = $1
	`

	var user model.User
	err := i.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Email,
		&user.Name,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return model.User{}, apperrors.NotFound("user not found")
	}

	if err != nil {
		return model.User{}, errors.Wrap(err, "failed to get user by email")
	}

	return user, nil
}
