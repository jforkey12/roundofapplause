package appService

import (
	m "dev/app/models"
	"net/url"
	"strconv"
	"strings"
)

func (svc appService) GetBugs(query url.Values) (bugs []m.Bug, err error) {
	filter := make(map[string]string)

	for key, value := range query {
		filter[key] = value[0]
	}

	testerIds, devices := ParseBugParams(filter)

	bugs, err = svc.db.GetBugs(testerIds, devices)
	return bugs, err
}

func ParseBugParams(stringFilter map[string]string) (testerIds []int, devices []string) {
	for key, _ := range stringFilter {
		if key == "createdBy" {
			testerIdsStr := strings.Split(stringFilter[key], ",")
			for _, idstr := range testerIdsStr {
				id, _ := strconv.Atoi(idstr)
				testerIds = append(testerIds, id)
			}
		}
		if key == "device" {
			devices = strings.Split(stringFilter[key], ",")
		}
	}
	return testerIds, devices
}
