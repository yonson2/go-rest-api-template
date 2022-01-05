package sql

import "github.com/yonson2/go-rest-api-template/pkg/models"

// CreateUser creates a new User
// returns the newly created user's ID.
func (s *DB) CreateUser(user *models.User) (string, error) {
	result := s.client.Create(&user)
	if result.Error != nil {
		return "", result.Error
	}

	return user.ID.String(), nil
}
