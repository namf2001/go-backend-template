package repository

import (
	"database/sql"

	"github.com/namf2001/go-backend-template/internal/repository/accounts"
	"github.com/namf2001/go-backend-template/internal/repository/sessions"
	"github.com/namf2001/go-backend-template/internal/repository/users"
)

type Registry interface {
	User() users.Repository
	Account() accounts.Repository
	Session() sessions.Repository
}

type impl struct {
	db       *sql.DB
	users    users.Repository
	accounts accounts.Repository
	sessions sessions.Repository
}

func (i impl) User() users.Repository {
	return i.users
}

func (i impl) Account() accounts.Repository {
	return i.accounts
}

func (i impl) Session() sessions.Repository {
	return i.sessions
}

func New(db *sql.DB) Registry {
	return &impl{
		db:       db,
		users:    users.New(db),
		accounts: accounts.New(db),
		sessions: sessions.New(db),
	}
}
