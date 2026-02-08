package accounts

import (
	"context"

	"github.com/namf2001/go-backend-template/internal/model"
	apperrors "github.com/namf2001/go-backend-template/internal/pkg/errors"
	"github.com/pkg/errors"
)

// Update updates an account.
func (i impl) Update(ctx context.Context, account model.Account) error {
	query := `
	UPDATE accounts
	SET username = $1, password = $2, updated_at = NOW()
	WHERE id = $3
	`

	result, err := i.db.ExecContext(ctx, query,
		account.Username,
		account.Password,
		account.ID,
	)
	if err != nil {
		if err.Error() == "pq: duplicate key value violates unique constraint \"users_email_key\"" ||
			err.Error() == "UNIQUE constraint failed" {
			return apperrors.AlreadyExists("account with this username already exists")
		}
		return errors.Wrap(err, "failed to update account")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "failed to get rows affected")
	}

	if rowsAffected == 0 {
		return apperrors.NotFound("account not found")
	}

	return nil

}
