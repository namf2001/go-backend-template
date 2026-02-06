package users

import (
	"context"

	"github.com/pkg/errors"
)

// CountUser returns the total number of users in the database.
func (i impl) CountUser(ctx context.Context) (int64, error) {
	query := `SELECT COUNT(*) FROM users`

	var count int64
	err := i.db.QueryRowContext(ctx, query).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "failed to count users")
	}

	return count, nil
}
