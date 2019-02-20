package bugsService

import (
	m "dev/bugs/models"
)

func (svc bugsService) CreateBug(bug m.Bug) (m.Bug, error) {
	bug, err := svc.db.InsertBug(bug)
	return bug, err
}
