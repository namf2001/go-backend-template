package users

import (
	"context"
	"testing"

	"github.com/namf2001/go-backend-template/internal/pkg/testdb"
	"github.com/namf2001/go-backend-template/internal/repository/db/pg"
	"github.com/stretchr/testify/require"
)

func TestDelete(t *testing.T) {
	type args struct {
		givenID int64
		expErr  error
	}

	tcs := map[string]args{
		"success": {
			givenID: 1001,
		},
		"err - user not found": {
			givenID: 99999,
			expErr:  ErrNotFound,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			testdb.WithTx(t, func(tx pg.ContextExecutor) {
				testdb.LoadTestSQLFile(t, tx, "testdata/users.sql")
				repo := New(tx)
				err := repo.Delete(context.Background(), tc.givenID)

				if tc.expErr != nil {
					require.ErrorIs(t, err, tc.expErr)
				} else {
					require.NoError(t, err)

					// Verify the user is actually deleted
					_, err = repo.GetByID(context.Background(), tc.givenID)
					require.Error(t, err)
				}
			})
		})
	}
}
