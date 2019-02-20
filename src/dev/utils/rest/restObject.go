package rest

import (
	"sync"
)

var once sync.Once
var instance RestInterface

type restObject struct{}

func GetRestInterface() RestInterface {
	once.Do(func() {
		instance = restObject{}
	})
	return instance
}
