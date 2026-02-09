package accounts

import (
	"context"
	"database/sql"

	"github.com/namf2001/go-backend-template/internal/model"
	apperrors "github.com/namf2001/go-backend-template/internal/pkg/errors"
	"github.com/pkg/errors"
)

// GetByProvider implements Repository.
func (i impl) GetByProvider(ctx context.Context, provider, providerAccountID string) (model.Account, error) {
	query := `
		SELECT id, "userId", type, provider, "providerAccountId", refresh_token, access_token, expires_at, id_token, scope, session_state, token_type
		FROM accounts
		WHERE provider = $1 AND "providerAccountId" = $2
	`

	var account model.Account
	err := i.db.QueryRowContext(ctx, query, provider, providerAccountID).Scan(
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
	)

	if err == sql.ErrNoRows {
		return model.Account{}, apperrors.NotFound("account not found")
	}

	if err != nil {
		return model.Account{}, errors.Wrap(err, "failed to get account by provider")
	}

	return account, nil
}
