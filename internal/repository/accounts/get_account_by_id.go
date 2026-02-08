package accounts

import (
	"context"
	"database/sql"

	"github.com/namf2001/go-backend-template/internal/model"
	apperrors "github.com/namf2001/go-backend-template/internal/pkg/errors"
	"github.com/pkg/errors"
)

// GetByID retrieves an account by ID.
func (i impl) GetByID(ctx context.Context, id int64) (model.Account, error) {
	query := `
	SELECT id, user_id, account_number, balance, created_at, updated_at
	FROM accounts
	WHERE id = $1
	`

	var account model.Account
	err := i.db.QueryRowContext(ctx, query, id).Scan(
		&account.ID,
		&account.UserID,
		&account.Username,
		&account.Password,
		&account.CreatedAt,
		&account.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return model.Account{}, apperrors.NotFound("account not found")
	}
	if err != nil {
		return model.Account{}, errors.Wrap(err, "failed to get account")
	}

	return account, nil
}
