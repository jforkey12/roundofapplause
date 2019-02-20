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
}
