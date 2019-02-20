package bugsService

import (
	m "dev/bugs/models"
)

func (svc bugsService) UpdateBug(id string, bug m.Bug) (m.Bug, error) {

	dbBug, err := svc.GetBugByID(id)
	if err != nil {
		return bug, err
	}

	dbBug.Merge(bug)

	bug, err = svc.db.ReplaceBug(dbBug)
	return bug, err
}
