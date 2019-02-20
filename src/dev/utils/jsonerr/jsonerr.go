package jsonerr

import (
	"encoding/json"
	"net/http"
)

type ErrJSON struct {
	Err string `json:"error"`
}

func JSONError(w http.ResponseWriter, err error, status int) {
	ErrorJSONObject := &ErrJSON{err.Error()}
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ErrorJSONObject)
}
