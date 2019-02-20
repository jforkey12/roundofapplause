package rest

import (
	"net/http"

	"github.com/go-resty/resty"
)

type RestInterface interface {
	CallRESTAPI(url string, method string, body interface{}, r *http.Request) (*resty.Response, error)
	HandleRESTAPIResponse(resty.Response, error, string) error
}
