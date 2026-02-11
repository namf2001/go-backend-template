package accounts

import (
	"context"

	"github.com/namf2001/go-backend-template/internal/model"
	pkgerrors "github.com/pkg/errors"
)

// GetByUserID implements Repository.
func (i impl) GetByUserID(ctx context.Context, userID int64) ([]model.Account, error) {
	query := `
		SELECT id, "userId", type, provider, "providerAccountId", refresh_token, access_token, expires_at, id_token, scope, session_state, token_type
		FROM accounts
		WHERE "userId" = $1
	`

	rows, err := i.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, pkgerrors.WithStack(err)
	}
	defer rows.Close()

	var accounts []model.Account
	for rows.Next() {
		var account model.Account
		if err := rows.Scan(
			&account.ID,
			&account.UserID,
			&account.Type,
			&account.Provider,
			&account.ProviderAccountID,
			&account.RefreshToken,
			&account.AccessToken,
			&account.ExpiresAt,
			&account.IDToken,
			&account.Scope,
			&account.SessionState,
			&account.TokenType,
		); err != nil {
			return nil, pkgerrors.WithStack(err)
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}
