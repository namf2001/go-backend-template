package users

import (
	"context"
	"testing"

	"github.com/namf2001/go-backend-template/internal/model"
	"github.com/namf2001/go-backend-template/internal/pkg/testdb"
	"github.com/namf2001/go-backend-template/internal/repository/db/pg"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	type args struct {
		givenUser model.User
		expErr    error
	}

	tcs := map[string]args{
		"success": {
			givenUser: model.User{
				Email:    "newuser@example.com",
				Name:     "New User",
				Password: "hashedpassword",
				Image:    "https://example.com/new.png",
			},
		},
		"err - duplicate email": {
			givenUser: model.User{
				Email:    "test1@example.com", // already exists in testdata
				Name:     "Duplicate User",
				Password: "hashedpassword",
			},
			expErr: ErrAlreadyExists,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			testdb.WithTx(t, func(tx pg.ContextExecutor) {
				testdb.LoadTestSQLFile(t, tx, "testdata/users.sql")
				repo := New(tx)
				created, err := repo.Create(context.Background(), tc.givenUser)

				if tc.expErr != nil {
					require.ErrorIs(t, err, tc.expErr)
				} else {
					require.NoError(t, err)
					require.NotZero(t, created.ID)
					require.Equal(t, tc.givenUser.Email, created.Email)
					require.Equal(t, tc.givenUser.Name, created.Name)
					require.NotZero(t, created.CreatedAt)
					require.NotZero(t, created.UpdatedAt)
				}
			})
		})
	}
}
