package usersService

import (
	m "dev/users/models"
)

func (svc usersService) GetUserByID(id string) (m.User, error) {
	object := m.User{}

	object, err := svc.db.GetUserByID(id)
	return object, err
}
