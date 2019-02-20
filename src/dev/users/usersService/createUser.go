package usersService

import (
	m "dev/users/models"
)

func (svc usersService) CreateUser(user m.User) (m.User, error) {
	user, err := svc.db.InsertUser(user)
	return user, err
}
