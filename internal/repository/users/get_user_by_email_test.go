package users

import (
	"context"
	"testing"

	"github.com/namf2001/go-backend-template/internal/model"
	"github.com/namf2001/go-backend-template/internal/pkg/testdb"
	"github.com/namf2001/go-backend-template/internal/repository/db/pg"
	"github.com/stretchr/testify/require"
)

func TestGetByEmail(t *testing.T) {
	type args struct {
		givenEmail string
		expErr     error
	}

	tcs := map[string]args{
		"success": {
			givenEmail: "test1@example.com",
		},
		"err - user not found": {
			givenEmail: "nonexistent@example.com",
			expErr:     model.ErrUserNotFound,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			testdb.WithTx(t, func(tx pg.ContextExecutor) {
				testdb.LoadTestSQLFile(t, tx, "testdata/users.sql")
				repo := New(tx)
				user, err := repo.GetByEmail(context.Background(), tc.givenEmail)

				if tc.expErr != nil {
					require.ErrorIs(t, err, tc.expErr)
				} else {
					require.NoError(t, err)
					require.Equal(t, tc.givenEmail, user.Email)
					require.NotZero(t, user.ID)
					require.NotEmpty(t, user.Name)
				}
			})
		})
	}
}
