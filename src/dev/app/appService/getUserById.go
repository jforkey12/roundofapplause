package appService

import (
	m "dev/app/models"
)

func (svc appService) GetUserByID(id string) (m.User, error) {
	object := m.User{}

	object, err := svc.db.GetUserByID(id)
	return object, err
}
