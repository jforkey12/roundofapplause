package bugsService

import (
	m "dev/bugs/models"
)

func (svc bugsService) GetBugByID(id string) (m.Bug, error) {
	object := m.Bug{}

	object, err := svc.db.GetBugByID(id)
	return object, err
}
