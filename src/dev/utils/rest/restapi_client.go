package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (restAccess restObject) CallRESTAPI(url string, method string, body interface{}) (*http.Response, error) {

	var data []byte
	if body != nil {
		bytes, err := json.Marshal(body)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(string(bytes))
		data, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	req, _ := http.NewRequest(method, url, bytes.NewBuffer(data))
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	req.Header.Add("cache-control", "no-cache")

	httpClient := &http.Client{}
	return httpClient.Do(req)
}
