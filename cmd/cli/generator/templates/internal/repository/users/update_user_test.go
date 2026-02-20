package users

import (
	"context"
	"testing"

	"github.com/namf2001/go-backend-template/internal/model"
	"github.com/namf2001/go-backend-template/internal/pkg/testdb"
	"github.com/namf2001/go-backend-template/internal/repository/db/pg"
	"github.com/stretchr/testify/require"
)

func TestUpdate(t *testing.T) {
	type args struct {
		givenUser model.User
		expErr    error
	}

	tcs := map[string]args{
		"success": {
			givenUser: model.User{
				ID:       1001,
				Email:    "updated@example.com",
				Name:     "Updated User",
				Password: "$2a$10$updatedpassword",
				Image:    "https://example.com/updated.png",
			},
		},
		"err - user not found": {
			givenUser: model.User{
				ID:       99999,
				Email:    "ghost@example.com",
				Name:     "Ghost User",
				Password: "hashedpassword",
			},
			expErr: ErrNotFound,
		},
		"err - duplicate email": {
			givenUser: model.User{
				ID:       1001,
				Email:    "test2@example.com", // belongs to user 1002
				Name:     "Conflict User",
				Password: "hashedpassword",
			},
			expErr: ErrDuplicateEmail,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			testdb.WithTx(t, func(tx pg.ContextExecutor) {
				testdb.LoadTestSQLFile(t, tx, "testdata/users.sql")
				repo := New(tx)
				err := repo.Update(context.Background(), tc.givenUser)

				if tc.expErr != nil {
					require.ErrorIs(t, err, tc.expErr)
				} else {
					require.NoError(t, err)

					// Verify the update took effect
					updated, err := repo.GetByID(context.Background(), tc.givenUser.ID)
					require.NoError(t, err)
					require.Equal(t, tc.givenUser.Email, updated.Email)
					require.Equal(t, tc.givenUser.Name, updated.Name)
				}
			})
		})
	}
}
