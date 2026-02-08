package users

import (
	"context"
	"strconv"

	"github.com/namf2001/go-backend-template/internal/model"
	"github.com/pkg/errors"
)

// ListFilters represents filters for listing users
type ListFilters struct {
	Limit  int
	Offset int
	Order  string
	Email  string
}

// List implements Repository.
func (i impl) List(ctx context.Context, filters ListFilters) ([]model.User, error) {
	query := `
		SELECT id, email, name, created_at, updated_at
		FROM users
		WHERE 1=1
	`
	var args []any
	argPos := 1

	// Add email filter if provided
	if filters.Email != "" {
		query += ` AND email ILIKE $` + strconv.Itoa(argPos)
		args = append(args, "%"+filters.Email+"%")
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
		return nil, errors.Wrap(err, "failed to list users")
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.Name,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, errors.Wrap(err, "failed to scan user")
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "error iterating users")
	}

	return users, nil
}
