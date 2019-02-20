package usersService

func (svc usersService) DeleteUser(id string) error {

	_, err := svc.db.GetUserByID(id)
	if err != nil {
		return err
	}
	err = svc.db.DeleteUser(id)
	return err
}
