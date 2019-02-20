package appService

func (svc appService) DeleteBug(id string) error {

	_, err := svc.db.GetBugByID(id)
	if err != nil {
		return err
	}
	err = svc.db.DeleteBug(id)
	return err
}
