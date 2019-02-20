package usersService

import (
	m "dev/users/models"
	"dev/users/usersDBAccess"
	"dev/utils/rest"

	"sync"
)

var once sync.Once
var instance m.ServiceInterface

type usersService struct {
	db      m.DbInterface
	restAdm rest.RestInterface
}

func GetService() m.ServiceInterface {
	once.Do(func() {
		rest := rest.GetRestInterface()
		dbInstance := usersDBAccess.GetMgoService()
		instance = usersService{db: dbInstance, restAdm: rest}
	})
	return instance
}
