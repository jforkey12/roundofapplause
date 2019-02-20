package handler

import (
	"dev/users/models"
	"dev/users/usersService"
	"dev/utils/rest"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

var once sync.Once
var instance models.RestInterface

type restService struct {
	service models.ServiceInterface
	restAdm rest.RestInterface
}

func getRestService() models.RestInterface {
	once.Do(func() {
		serv := usersService.GetService()
		rest := rest.GetRestInterface()
		instance = restService{restAdm: rest, service: serv}
	})
	return instance
}

type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type routes []route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	restServer := getRestService()
	restServer.Init()

	r := routes{}
	r = addUsersRoutes(restServer, r)

	for _, route := range r {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = rest.RestAPIEntryPoint(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		_, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		return nil
	})

	return router
}

func addUsersRoutes(rest models.RestInterface, r routes) routes {
	r = append(r,
		route{
			"CreateUser",
			"POST",
			"/applause/v1/users",
			rest.CreateUser,
		},
		route{
			"GetUsers",
			"GET",
			"/applause/v1/users",
			rest.GetUsers,
		},
		route{
			"GetUserByID",
			"GET",
			"/applause/v1/users/{id}",
			rest.GetUserByID,
		},
		route{
			"UpdateUser",
			"PATCH",
			"/applause/v1/users/{id}",
			rest.UpdateUser,
		},
		route{
			"DeleteUser",
			"DELETE",
			"/applause/v1/users/{id}",
			rest.DeleteUser,
		},
	)
	return r
}
