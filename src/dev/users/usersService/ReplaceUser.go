package usersService

import (
	m "dev/users/models"
)

func (svc usersService) ReplaceUser(id string, user m.User) (m.User, error) {

	dbUser, err := svc.GetUserByID(id)
	if err != nil {
		return user, err
	}

	dbUser.Merge(user)

	userInfo, err := svc.db.ReplaceUser(user)
	return userInfo, err
}
