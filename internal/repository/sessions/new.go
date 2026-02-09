package sessions

import (
	"context"
	"database/sql"

	"github.com/namf2001/go-backend-template/internal/model"
)

type Repository interface {
	// Create creates a new session
	Create(ctx context.Context, session model.Session) (model.Session, error)

	// GetByToken retrieves a session by session token
	GetByToken(ctx context.Context, token string) (model.Session, error)

	// Delete deletes a session by session token
	Delete(ctx context.Context, token string) error
}

type impl struct {
	db *sql.DB
}

func New(db *sql.DB) Repository {
	return impl{
		db: db,
	}
}
