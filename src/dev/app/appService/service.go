package appService

import (
	"dev/app/dbAccess"
	m "dev/app/models"

	"sync"
)

var once sync.Once
var instance m.ServiceInterface

type appService struct {
	db m.DbInterface
}

func GetService() m.ServiceInterface {
	once.Do(func() {
		dbInstance := dbAccess.GetMgoService()
		instance = appService{db: dbInstance}
	})
	return instance
}
