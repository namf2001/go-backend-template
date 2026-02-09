package sessions

import (
	"context"

	"github.com/namf2001/go-backend-template/internal/model"
	"github.com/pkg/errors"
)

// Create implements Repository.
func (i impl) Create(ctx context.Context, session model.Session) (model.Session, error) {
	query := `
		INSERT INTO sessions ("userId", expires, "sessionToken")
		VALUES ($1, $2, $3)
		RETURNING id, "userId", expires, "sessionToken"
	`

	var created model.Session
	err := i.db.QueryRowContext(ctx, query, session.UserID, session.Expires, session.SessionToken).Scan(
		&created.ID,
		&created.UserID,
		&created.Expires,
		&created.SessionToken,
	)

	if err != nil {
		return model.Session{}, errors.Wrap(err, "failed to create session")
	}

	return created, nil
}
