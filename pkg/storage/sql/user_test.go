package sql_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/yonson2/go-rest-api-template/pkg/models"
	"github.com/yonson2/go-rest-api-template/pkg/storage/sql"
)

func TestTeamService_CreateTeam(t *testing.T) {
	is := is.New(t)

	storage, err := sql.NewSqliteDB("file::memory:?cache=shared")
	is.NoErr(err)
	storage.Migrate()
	_, err = storage.CreateUser(&models.User{Username: "Bonobxs", Email: "test@example.org"})
	is.NoErr(err)
}
