package models

type DbInterface interface {
	InitSession() error

	InsertBug(Bug) (Bug, error)
	ReplaceBug(Bug) (Bug, error)
	GetBugs(testers []int, devices []string) ([]Bug, error)
	GetBugByID(string) (Bug, error)
	DeleteBug(string) error

	InsertUser(User) (User, error)
	ReplaceUser(User) (User, error)
	GetUsers(countries []string, devices []string) ([]User, error)
	GetUserByID(string) (User, error)
	DeleteUser(string) error
}
