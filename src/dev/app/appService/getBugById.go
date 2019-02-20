package appService

import (
	m "dev/app/models"
)

func (svc appService) GetBugByID(id string) (m.Bug, error) {
	object := m.Bug{}

	object, err := svc.db.GetBugByID(id)
	return object, err
}
