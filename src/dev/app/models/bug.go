package models

type Bug struct {
	ID        int    `json:"id"`
	Device    string `json:"device"`
	CreatedBy int    `json:"createdBy"`
}

func (a *Bug) Merge(a2 Bug) {
	if a2.Device != "" {
		a.Device = a2.Device
	}
	if a2.CreatedBy != 0 {
		a.CreatedBy = a2.CreatedBy
	}
}
