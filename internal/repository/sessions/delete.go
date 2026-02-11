package sessions

import (
	"context"

	pkgerrors "github.com/pkg/errors"
)

// Delete implements Repository.
func (i impl) Delete(ctx context.Context, token string) error {
	query := `
		DELETE FROM sessions
		WHERE "sessionToken" = $1
	`

	_, err := i.db.ExecContext(ctx, query, token)
	if err != nil {
		return pkgerrors.WithStack(err)
	}

	return nil
}
