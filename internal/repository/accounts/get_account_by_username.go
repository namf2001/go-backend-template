package accounts

import (
	"context"
	"database/sql"

	"github.com/namf2001/go-backend-template/internal/model"
	apperrors "github.com/namf2001/go-backend-template/internal/pkg/errors"
	"github.com/pkg/errors"
)

// GetByUserName retrieves an account by username.
func (i impl) GetByUserName(ctx context.Context, email string) (model.Account, error) {
	query := `
	SELECT id, user_id, account_number, balance, created_at, updated_at
	FROM accounts
	WHERE username = $1
	`

	var account model.Account
	err := i.db.QueryRowContext(ctx, query, email).Scan(
		&account.ID,
		&account.UserID,
		&account.Username,
		&account.Password,
		&account.CreatedAt,
		&account.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return model.Account{}, apperrors.NotFound("user not found")
	}

	if err != nil {
		return model.Account{}, errors.Wrap(err, "failed to get account by username")
	}

	return account, nil
}
