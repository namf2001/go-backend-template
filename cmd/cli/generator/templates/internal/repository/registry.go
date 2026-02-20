package repository

import (
	"context"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/namf2001/go-backend-template/internal/repository/accounts"
	"github.com/namf2001/go-backend-template/internal/repository/db/pg"
	"github.com/namf2001/go-backend-template/internal/repository/sessions"
	"github.com/namf2001/go-backend-template/internal/repository/users"
	pkgerrors "github.com/pkg/errors"
)

// Registry is the registry of all the domain specific repositories and also provides transaction capabilities.
type Registry interface {
	// User return user repository
	User() users.Repository
	// Account return account repository
	Account() accounts.Repository
	// Session return session repository
	Session() sessions.Repository
	// DoInTx wraps operations within a db tx
	DoInTx(ctx context.Context, txFunc func(ctx context.Context, txRepo Registry) error, overrideBackoffPolicy backoff.BackOff) error
}

// New returns a new instance of Registry
func New(db pg.BeginnerExecutor) Registry {
	return &impl{
		pgConn:   db,
		users:    users.New(db),
		accounts: accounts.New(db),
		sessions: sessions.New(db),
	}
}

type impl struct {
	pgConn   pg.BeginnerExecutor // Only used to start DB txns
	tx       pg.ContextExecutor  // Only used to keep track if txn has already been started to prevent nested txns
	users    users.Repository
	accounts accounts.Repository
	sessions sessions.Repository
}

func (i *impl) User() users.Repository {
	return i.users
}

func (i *impl) Account() accounts.Repository {
	return i.accounts
}

func (i *impl) Session() sessions.Repository {
	return i.sessions
}

// DoInTx wraps operations within a db tx.
// It creates a new Registry where all repositories share the same transaction.
// Nested transactions are not allowed.
func (i *impl) DoInTx(ctx context.Context, txFunc func(ctx context.Context, txRepo Registry) error, overrideBackoffPolicy backoff.BackOff) error {
	if i.tx != nil {
		return pkgerrors.WithStack(errNestedTx)
	}

	if overrideBackoffPolicy == nil {
		overrideBackoffPolicy = pg.ExponentialBackOff(3, time.Minute)
	}

	return pg.TxWithBackOff(ctx, overrideBackoffPolicy, i.pgConn, func(tx pg.ContextExecutor) error {
		newI := &impl{
			tx:       tx,
			users:    users.New(tx),
			accounts: accounts.New(tx),
			sessions: sessions.New(tx),
		}
		return txFunc(ctx, newI)
	})
}
