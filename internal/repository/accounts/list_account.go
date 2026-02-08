package accounts

import (
	"context"
	"strconv"

	"github.com/namf2001/go-backend-template/internal/model"
	"github.com/pkg/errors"
)

// ListFilters defines filters for listing accounts
type ListFilters struct {
	Limit    int
	Offset   int
	Order    string
	Username string
}

// List retrieves accounts with optional filters
func (i impl) List(ctx context.Context, filters ListFilters) ([]model.Account, error) {
	query := `
		SELECT id, user_name, email, created_at, updated_at
		FROM accounts
		WHERE 1=1
	`
	var args []any
	argPos := 1

	// Add user_name filter if provided
	if filters.Username != "" {
		query += ` AND user_name ILIKE $` + strconv.Itoa(argPos)
		args = append(args, "%"+filters.Username+"%")
		argPos++
	}

	// Add ordering
	query += ` ORDER BY created_at `
	if filters.Order == "asc" {
		query += `ASC`
	} else {
		query += `DESC`
	}

	// Add pagination
	if filters.Limit > 0 {
		query += ` LIMIT $` + strconv.Itoa(argPos)
		args = append(args, filters.Limit)
		argPos++
	}
	if filters.Offset > 0 {
		query += ` OFFSET $` + strconv.Itoa(argPos)
		args = append(args, filters.Offset)
	}

	rows, err := i.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to list accounts")
	}
	defer rows.Close()

	var accounts []model.Account
	for rows.Next() {
		var account model.Account
		err := rows.Scan(
			&account.ID,
			&account.Username,
			&account.CreatedAt,
			&account.UpdatedAt,
		)
		if err != nil {
			return nil, errors.Wrap(err, "failed to scan account")
		}
		accounts = append(accounts, account)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "error iterating accounts")
	}

	return accounts, nil
}
