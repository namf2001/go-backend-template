package users

import (
	"context"

	apperrors "github.com/namf2001/go-backend-template/internal/pkg/errors"
	"github.com/pkg/errors"
)

// Delete implements Repository.
func (i impl) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM users WHERE id = $1`

	result, err := i.db.ExecContext(ctx, query, id)
	if err != nil {
		return errors.Wrap(err, "failed to delete user")
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
