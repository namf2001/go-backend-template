package users

import (
	"context"
	"database/sql"

	"github.com/namf2001/go-backend-template/internal/model"
	apperrors "github.com/namf2001/go-backend-template/internal/pkg/errors"
	"github.com/pkg/errors"
)

// GetByID implements Repository.
func (i impl) GetByID(ctx context.Context, id int64) (model.User, error) {
	query := `
		SELECT id, email, name, password, image, "emailVerified", created_at, updated_at
		FROM users
		WHERE id = $1
	`

	var user model.User
	err := i.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Email,
		&user.Name,
		&user.Password,
		&user.Image,
		&user.EmailVerified,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return model.User{}, apperrors.NotFound("user not found")
	}
	if err != nil {
		return model.User{}, errors.Wrap(err, "failed to get user")
	}

	return user, nil
}
