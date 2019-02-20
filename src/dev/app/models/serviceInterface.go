package models

import "net/url"

type ServiceInterface interface {
	CreateBug(Bug) (Bug, error)
	UpdateBug(string, Bug) (Bug, error)
	ReplaceBug(string, Bug) (Bug, error)
	GetBugs(query url.Values) ([]Bug, error)
	GetBugByID(string) (Bug, error)
	DeleteBug(string) error

	CreateUser(User) (User, error)
	UpdateUser(string, User) (User, error)
	ReplaceUser(string, User) (User, error)
	GetUsers(query url.Values) ([]User, error)
	GetUserByID(string) (User, error)
	DeleteUser(string) error
}
