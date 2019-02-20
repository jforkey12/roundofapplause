package appService

import (
	m "dev/app/models"
)

func (svc appService) CreateBug(bug m.Bug) (m.Bug, error) {
	bug, err := svc.db.InsertBug(bug)
	return bug, err
}
