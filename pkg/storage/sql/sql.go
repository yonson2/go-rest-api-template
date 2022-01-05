package sql

import (
	"github.com/yonson2/go-rest-api-template/pkg/models"
	"github.com/yonson2/go-rest-api-template/pkg/storage"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Ensure service implements interface.
var _ storage.Storage = (*DB)(nil)

type DB struct {
	client *gorm.DB
}

// NewSqliteDB returns a Db client, backed by an sqlite driver.
func NewSqliteDB(dsn string) (*DB, error) {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &DB{client: db}, nil
}

func (sql *DB) Migrate() error {
	err := sql.client.AutoMigrate(&models.User{})

	return err
}
