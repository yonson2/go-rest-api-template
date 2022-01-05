package sql_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/yonson2/go-rest-api-template/pkg/models"
	"github.com/yonson2/go-rest-api-template/pkg/storage/sql"
)

func TestTeamService_CreateTeam(t *testing.T) {
	storage, err := sql.NewSqliteDB("file::memory:?cache=shared")

	t.Run("it creates a user", func(t *testing.T) {
		t.Parallel()
		is := is.New(t)

		is.NoErr(err)

		_ = storage.Migrate()
		_, err = storage.CreateUser(&models.User{Username: "Bonobxs", Email: "test@example.org"})
		is.NoErr(err)
	})
}
