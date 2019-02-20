package models

type User struct {
	ID        int      `json:"id"`
	FirstName string   `json:"firstName"`
	LastName  string   `json:"lastName"`
	Country   string   `json:"country"`
	LastLogin string   `json:"lastLogin"`
	Devices   []string `json:"devices"`
	BugCount  int      `json:"bugCount"`
	Bugs      []string `json:"bugs"`
}

func (a *User) Merge(a2 User) {
	if a2.FirstName != "" {
		a.FirstName = a2.FirstName
	}
	if a2.LastName != "" {
		a.LastName = a2.LastName
	}
	if a2.Country != "" {
		a.Country = a2.Country
	}
	if len(a2.Devices) > 0 {
		a.Devices = a2.Devices
	}

}
