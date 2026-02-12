package users

import (
	"context"
	"errors"

	"github.com/lib/pq"
	"github.com/namf2001/go-backend-template/internal/model"
	pkgerrors "github.com/pkg/errors"
)

// Create implements Repository.
func (i impl) Create(ctx context.Context, user model.User) (model.User, error) {
	query := `
		INSERT INTO users (email, name, password, image, "emailVerified")
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, email, name, password, image, "emailVerified", created_at, updated_at
	`

	var created model.User
	err := i.db.QueryRowContext(ctx, query, user.Email, user.Name, user.Password, user.Image, user.EmailVerified).Scan(
		&created.ID,
		&created.Email,
		&created.Name,
		&created.Password,
		&created.Image,
		&created.EmailVerified,
		&created.CreatedAt,
		&created.UpdatedAt,
	)

	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == "23505" {
			return model.User{}, pkgerrors.WithStack(ErrAlreadyExists)
		}

		return model.User{}, pkgerrors.WithStack(err)
	}

	return created, nil
}
