package repository

import (
	"database/sql"

	"github.com/namf2001/go-backend-template/internal/repository/users"
)

type Registry interface {
	User() users.Repository
}

type impl struct {
	db    *sql.DB
	users users.Repository
}

func (i impl) User() users.Repository {
	return i.users
}

func New(db *sql.DB) Registry {
	return &impl{
		db:    db,
		users: users.New(db),
	}
}
