package models

import (
	"net/http"
)

type RestInterface interface {
	Init() error
	CreateBug(w http.ResponseWriter, r *http.Request)
	GetBugs(w http.ResponseWriter, r *http.Request)
	GetBugByID(w http.ResponseWriter, r *http.Request)
	UpdateBug(w http.ResponseWriter, r *http.Request)
	ReplaceBug(w http.ResponseWriter, r *http.Request)
	DeleteBug(w http.ResponseWriter, r *http.Request)

	CreateUser(w http.ResponseWriter, r *http.Request)
	GetUsers(w http.ResponseWriter, r *http.Request)
	GetUserByID(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	ReplaceUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}
