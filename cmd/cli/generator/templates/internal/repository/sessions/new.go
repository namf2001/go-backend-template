package sessions

import (
	"context"

	"github.com/namf2001/go-backend-template/internal/model"
	"github.com/namf2001/go-backend-template/internal/repository/db/pg"
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
	db pg.ContextExecutor
}

func New(db pg.ContextExecutor) Repository {
	return impl{
		db: db,
	}
}
