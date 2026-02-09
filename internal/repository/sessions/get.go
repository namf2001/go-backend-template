package sessions

import (
	"context"
	"database/sql"

	"github.com/namf2001/go-backend-template/internal/model"
	apperrors "github.com/namf2001/go-backend-template/internal/pkg/errors"
	"github.com/pkg/errors"
)

// GetByToken implements Repository.
func (i impl) GetByToken(ctx context.Context, token string) (model.Session, error) {
	query := `
		SELECT id, "userId", expires, "sessionToken"
		FROM sessions
		WHERE "sessionToken" = $1
	`

	var session model.Session
	err := i.db.QueryRowContext(ctx, query, token).Scan(
		&session.ID,
		&session.UserID,
		&session.Expires,
		&session.SessionToken,
	)

	if err == sql.ErrNoRows {
		return model.Session{}, apperrors.NotFound("session not found")
	}

	if err != nil {
		return model.Session{}, errors.Wrap(err, "failed to get session by token")
	}

	return session, nil
}
