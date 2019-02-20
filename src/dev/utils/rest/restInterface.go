package rest

import (
	"net/http"
)

type RestInterface interface {
	CallRESTAPI(url string, method string, body interface{}) (*http.Response, error)
}
