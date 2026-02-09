package accounts

import (
	"context"

	"github.com/namf2001/go-backend-template/internal/model"
	"github.com/pkg/errors"
)

// Create implements Repository.
func (i impl) Create(ctx context.Context, account model.Account) (model.Account, error) {
	query := `
		INSERT INTO accounts ("userId", type, provider, "providerAccountId", refresh_token, access_token, expires_at, id_token, scope, session_state, token_type)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING id, "userId", type, provider, "providerAccountId", refresh_token, access_token, expires_at, id_token, scope, session_state, token_type
	`

	var created model.Account
	err := i.db.QueryRowContext(ctx, query,
		account.UserID,
		account.Type,
		account.Provider,
		account.ProviderAccountID,
		account.RefreshToken,
		account.AccessToken,
		account.ExpiresAt,
		account.IDToken,
		account.Scope,
		account.SessionState,
		account.TokenType,
	).Scan(
		&created.ID,
		&created.UserID,
		&created.Type,
		&created.Provider,
		&created.ProviderAccountID,
		&created.RefreshToken,
		&created.AccessToken,
		&created.ExpiresAt,
		&created.IDToken,
		&created.Scope,
		&created.SessionState,
		&created.TokenType,
	)

	if err != nil {
		return model.Account{}, errors.Wrap(err, "failed to create account")
	}

	return created, nil
}
