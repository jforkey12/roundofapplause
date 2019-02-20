package usersService

import (
	m "dev/users/models"
)

func (svc usersService) UpdateUser(id string, user m.User) (m.User, error) {

	dbUser, err := svc.GetUserByID(id)
	if err != nil {
		return user, err
	}

	dbUser.Merge(user)

	user, err = svc.db.ReplaceUser(dbUser)
	return user, err
}
