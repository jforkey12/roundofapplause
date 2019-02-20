package appService

import (
	m "dev/app/models"
)

func (svc appService) ReplaceBug(id string, bug m.Bug) (m.Bug, error) {

	dbBug, err := svc.GetBugByID(id)
	if err != nil {
		return bug, err
	}

	dbBug.Merge(bug)

	bugInfo, err := svc.db.ReplaceBug(bug)
	return bugInfo, err
}
