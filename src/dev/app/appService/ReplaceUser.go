package appService

import (
	m "dev/app/models"
)

func (svc appService) ReplaceUser(id string, user m.User) (m.User, error) {

	dbUser, err := svc.GetUserByID(id)
	if err != nil {
		return user, err
	}

	dbUser.Merge(user)

	userInfo, err := svc.db.ReplaceUser(user)
	return userInfo, err
}
