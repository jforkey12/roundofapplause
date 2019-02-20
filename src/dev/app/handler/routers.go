package handler

import (
	"dev/app/appService"
	"dev/app/models"
	"dev/utils/rest"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

var once sync.Once
var instance models.RestInterface

type restService struct {
	service models.ServiceInterface
}

func getRestService() models.RestInterface {
	once.Do(func() {
		serv := appService.GetService()
		instance = restService{service: serv}
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
	r = addRoutes(restServer, r)

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

func addRoutes(rest models.RestInterface, r routes) routes {
	r = append(r,
		route{
			"CreateBug",
			"POST",
			"/applause/v1/bugs",
			rest.CreateBug,
		},
		route{
			"GetBugs",
			"GET",
			"/applause/v1/bugs",
			rest.GetBugs,
		},
		route{
			"GetAppByID",
			"GET",
			"/applause/v1/bugs/{id}",
			rest.GetBugByID,
		},
		route{
			"UpdateBug",
			"PATCH",
			"/applause/v1/bugs/{id}",
			rest.UpdateBug,
		},
		route{
			"ReplaceBug",
			"PUT",
			"/applause/v1/bugs/{id}",
			rest.ReplaceBug,
		},
		route{
			"DeleteBug",
			"Delete",
			"/applause/v1/bugs/{id}",
			rest.DeleteBug,
		},
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
