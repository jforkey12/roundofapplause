package appService

import (
	m "dev/app/models"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

func (svc appService) GetUsers(query url.Values) (users []m.User, err error) {
	filter := make(map[string]string)

	for key, value := range query {
		filter[key] = value[0]
	}

	countries, devices := ParseUserParams(filter)
	users, err = svc.db.GetUsers(countries, devices)
	fmt.Println(strconv.Itoa(len(users)) + "found for inputs")
	var ids []int

	for _, user := range users {
		var ids []int
		ids = append(ids, user.ID)
	}

	bugs, _ := svc.db.GetBugs(ids, devices)

	for i, user := range users {
		bugCount := 0
		for _, bug := range bugs {
			if bug.CreatedBy == user.ID {
				bugCount++
			}
		}
		users[i].BugCount = bugCount
	}

	return users, err
}

func ParseUserParams(stringFilter map[string]string) (countries []string, devices []string) {
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
