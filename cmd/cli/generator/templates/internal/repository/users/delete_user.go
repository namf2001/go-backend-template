package users

import (
	"context"

	pkgerrors "github.com/pkg/errors"
)

// Delete deletes a user by ID.
func (i impl) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM users WHERE id = $1`

	result, err := i.db.ExecContext(ctx, query, id)
	if err != nil {
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
