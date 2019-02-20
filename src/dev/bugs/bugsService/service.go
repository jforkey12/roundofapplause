package bugsService

import (
	"dev/bugs/bugsDBAccess"
	m "dev/bugs/models"
	"dev/utils/rest"

	"sync"
)

var once sync.Once
var instance m.ServiceInterface

type bugsService struct {
	db      m.DbInterface
	restAdm rest.RestInterface
}

func GetService() m.ServiceInterface {
	once.Do(func() {
		rest := rest.GetRestInterface()
		dbInstance := bugsDBAccess.GetMgoService()
		instance = bugsService{db: dbInstance, restAdm: rest}
	})
	return instance
}
