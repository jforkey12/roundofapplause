package appService

import (
	m "dev/app/models"
)

func (svc appService) UpdateUser(id string, user m.User) (m.User, error) {

	dbUser, err := svc.GetUserByID(id)
	if err != nil {
		return user, err
	}

	dbUser.Merge(user)

	user, err = svc.db.ReplaceUser(dbUser)
	return user, err
}
