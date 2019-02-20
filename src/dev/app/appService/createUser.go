package appService

import (
	m "dev/app/models"
)

func (svc appService) CreateUser(user m.User) (m.User, error) {
	user, err := svc.db.InsertUser(user)
	return user, err
}
