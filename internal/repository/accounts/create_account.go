package accounts

import (
	"context"

	"github.com/namf2001/go-backend-template/internal/model"
	"github.com/pkg/errors"
)

// Create creates a new account.
func (i impl) Create(ctx context.Context, account model.Account) (model.Account, error) {
	query := `
		INSERT INTO accounts (userID, username, password, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	err := i.db.QueryRowContext(
		ctx,
		query,
		account.UserID,
		account.Username,
		account.Password,
		account.CreatedAt,
		account.UpdatedAt,
	).Scan(&account.ID)
	if err != nil {
		if err.Error() == "pq: duplicate key value violates unique constraint \"accounts_username_key\"" ||
			err.Error() == "UNIQUE constraint failed" {

			return model.Account{}, errors.New("account with this username already exists")
		}

		return model.Account{}, errors.Wrap(err, "failed to create user")
	}

	return account, nil
}
