package rest

import (
	"dev/utils/jsonerr"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/go-resty/resty"
)

func (rest restObject) CallRESTAPI(path string, method string, body interface{}, r *http.Request) (*resty.Response, error) {

	rClient := resty.New()
	rClient.SetLogger(ioutil.Discard)
	rClient.SetCloseConnection(true)

	request := rClient.R()
	request.Header.Add("Accept", "application/json; charset=UTF-8")
	request.Header.Add("Connection", "close")

	switch strings.ToUpper(method) {
	case "GET":
		if r != nil {
			if len(r.URL.RawQuery) != 0 {
				path = path + "?" + r.URL.RawQuery
			}
		}
		response, err := request.Get(path)
		return response, err
	case "POST":
		response, err := request.Post(path)
		return response, err
	case "PUT":
		response, err := request.Put(path)
		return response, err
	case "PATCH":
		response, err := request.Patch(path)
		return response, err
	case "DELETE":
		response, err := request.Delete(path)
		return response, err

		return nil, fmt.Errorf("Invalid method %v", method)
	}
	return nil, fmt.Errorf("Invalid method %v", method)
}

func (rest restObject) HandleRESTAPIResponse(response resty.Response, err error, errMsg string) error {
	if response.StatusCode() >= 400 {
		var jsonErr jsonerr.ErrJSON
		json.Unmarshal(response.Body(), &jsonErr)
		if errMsg != "" {
			errMsg += " " + jsonErr.Err
		} else {
			errMsg = jsonErr.Err
		}
		return errors.New(errMsg)
	}
	return nil

}
