package usersService

import (
	b "dev/bugs/models"
	m "dev/users/models"
	"encoding/json"
	"fmt"
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
		query := "device" + strings.Join(devices, ",") + "&id=" + strconv.Itoa(user.ID)
		response, err := svc.restAdm.CallRESTAPI("http://localhost:555/applause/v1/bugs?"+query, "GET", nil, nil)
		if err != nil {
			fmt.Println(response)
			fmt.Println(err)
			return users, err
		}

		err = svc.restAdm.HandleRESTAPIResponse(*response, err, "")
		if err != nil {
			fmt.Println(err)
			return users, err
		}

		var bugs []b.Bug
		err = json.Unmarshal([]byte(response.Body()), &bugs)
		if err != nil {
			fmt.Println(err)
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
