package appService

import (
	m "dev/app/models"
	"fmt"
	"net/url"
	"sort"
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

	sort.Slice(bugs, func(a, b int) bool {
		return bugs[a].CreatedBy < bugs[b].CreatedBy
	})
	fmt.Println(users)
	fmt.Println(bugs)
	j := 0
	for i, _ := range users {
		bugCount := 0
		for ; j <= len(bugs)-1; j++ {
			if bugs[j].CreatedBy == users[i].ID {
				bugCount++
				users[i].BugCount = bugCount
			}
			if bugs[j].CreatedBy > users[i].ID {
				break
			}
		}
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
