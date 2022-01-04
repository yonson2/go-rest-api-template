package storage

import (
	"github.com/yonson2/go-rest-api-template/pkg/models"
)

type Storage interface {
	CreateUser(user *models.User) (string, error)
}
