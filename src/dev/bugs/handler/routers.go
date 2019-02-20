package handler

import (
	"dev/bugs/bugsService"
	"dev/bugs/models"
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
		serv := bugsService.GetService()
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
	r = addBugsRoutes(restServer, r)

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

func addBugsRoutes(rest models.RestInterface, r routes) routes {
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
	)
	return r
}
