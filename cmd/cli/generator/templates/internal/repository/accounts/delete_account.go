package accounts

import (
	"context"

	pkgerrors "github.com/pkg/errors"
)

// Delete implements Repository.
func (i impl) Delete(ctx context.Context, provider, providerAccountID string) error {
	query := `
		DELETE FROM accounts
		WHERE provider = $1 AND "providerAccountId" = $2
	`

	_, err := i.db.ExecContext(ctx, query, provider, providerAccountID)
	if err != nil {
		return pkgerrors.WithStack(err)
	}

	return nil
}
