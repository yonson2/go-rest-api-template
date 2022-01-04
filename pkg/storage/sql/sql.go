package sql

import (
	"github.com/yonson2/go-rest-api-template/pkg/models"
	"github.com/yonson2/go-rest-api-template/pkg/storage"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Ensure service implements interface
var _ storage.Storage = (*SqlDB)(nil)

type SqlDB struct {
	client *gorm.DB
}

// NewSqliteDB returns an SqlDB client, backed by an sqlite driver
func NewSqliteDB(dsn string) (*SqlDB, error) {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &SqlDB{client: db}, nil
}

func (sql *SqlDB) Migrate() {
	sql.client.AutoMigrate(&models.User{})
}
