package accounts

import (
	"context"

	"github.com/pkg/errors"
)

// CountAccount returns the total number of accounts
func (i impl) CountAccount(ctx context.Context) (int64, error) {
	query := `SELECT COUNT(*) FROM accounts`

	var count int64
	err := i.db.QueryRowContext(ctx, query).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "failed to count users")
	}
	return count, nil
}
