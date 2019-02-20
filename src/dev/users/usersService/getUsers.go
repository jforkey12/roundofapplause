package usersService

import (
	b "dev/bugs/models"
	m "dev/users/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strconv"
	"strings"
)

func (svc usersService) GetUsers(query url.Values) (users []m.User, err error) {
	filter := make(map[string]string)

	for key, value := range query {
		filter[key] = value[0]
	}

	countries, devices := ParseQueryParams(filter)
	users, err = svc.db.GetUsers(countries, devices)
	for i, user := range users {
		query := "createdBy=" + strconv.Itoa(user.ID)

		if len(devices) > 0 {
			query = query + "&device=" + strings.Join(devices, ",")
		}
		equery := &url.URL{Path: query}
		path := "http://localhost:555/applause/v1/bugs?" + equery.String()

		response, err := svc.restAdm.CallRESTAPI(path, "GET", nil)
		if err != nil {
			fmt.Println(response)
			fmt.Println(err)
			return users, err
		}

		defer response.Body.Close()
		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
		}
		bugs := []b.Bug{}

		err = json.Unmarshal(data, &bugs)
		if err != nil {
			fmt.Println("unmarshal error:" + err.Error())
			return users, err
		}
		users[i].BugCount = len(bugs)
	}
	return users, err
}

func ParseQueryParams(stringFilter map[string]string) (countries []string, devices []string) {
	for key, _ := range stringFilter {
		if key == "country" {
			countries = strings.Split(stringFilter[key], ",")
		}
		if key == "devices" {
			devices = strings.Split(stringFilter[key], ",")
		}
	}
	return countries, devices
}
