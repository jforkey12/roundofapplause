package models

import "net/url"

type ServiceInterface interface {
	CreateUser(User) (User, error)
	UpdateUser(string, User) (User, error)
	ReplaceUser(string, User) (User, error)
	GetUsers(query url.Values) ([]User, error)
	GetUserByID(string) (User, error)
	DeleteUser(string) error
}
