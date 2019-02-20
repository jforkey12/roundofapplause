package models

type DbInterface interface {
	InitSession() error
	InsertUser(User) (User, error)
	ReplaceUser(User) (User, error)
	GetUsers(countries []string, devices []string) ([]User, error)
	GetUserByID(string) (User, error)
	DeleteUser(string) error
}
