package models

import "net/url"

type ServiceInterface interface {
	CreateBug(Bug) (Bug, error)
	UpdateBug(string, Bug) (Bug, error)
	ReplaceBug(string, Bug) (Bug, error)
	GetBugs(query url.Values) ([]Bug, error)
	GetBugByID(string) (Bug, error)
	DeleteBug(string) error
}
