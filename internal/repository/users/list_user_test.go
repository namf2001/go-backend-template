package users

import (
	"context"
	"testing"

	"github.com/namf2001/go-backend-template/internal/pkg/testdb"
	"github.com/namf2001/go-backend-template/internal/repository/db/pg"
	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	type args struct {
		givenFilters ListFilters
		expLen       int
	}

	tcs := map[string]args{
		"success - no filters": {
			givenFilters: ListFilters{},
			expLen:       3,
		},
		"success - with limit": {
			givenFilters: ListFilters{
				Limit: 2,
			},
			expLen: 2,
		},
		"success - with offset": {
			givenFilters: ListFilters{
				Limit:  10,
				Offset: 1,
			},
			expLen: 2,
		},
		"success - filter by email": {
			givenFilters: ListFilters{
				Email: "admin",
			},
			expLen: 1,
		},
		"success - filter by email no match": {
			givenFilters: ListFilters{
				Email: "nonexistent",
			},
			expLen: 0,
		},
		"success - order asc": {
			givenFilters: ListFilters{
				Order: "asc",
			},
			expLen: 3,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			testdb.WithTx(t, func(tx pg.ContextExecutor) {
				testdb.LoadTestSQLFile(t, tx, "testdata/users.sql")
				repo := New(tx)
				users, err := repo.List(context.Background(), tc.givenFilters)

				require.NoError(t, err)
				require.Len(t, users, tc.expLen)

				// Verify all returned users have required fields
				for _, u := range users {
					require.NotZero(t, u.ID)
					require.NotEmpty(t, u.Email)
					require.NotEmpty(t, u.Name)
				}
			})
		})
	}
}
