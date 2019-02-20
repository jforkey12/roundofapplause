package handler

import (
	m "dev/users/models"
	"dev/utils/jsonerr"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func (rest restService) Init() error {
	persistData := true
	if persistData {
		var devList []string
		devcsv := "data/devices.csv"

		f1, err := os.Open(devcsv)
		if err != nil {
			panic(err)
		}
		defer f1.Close()

		dLines, err := csv.NewReader(f1).ReadAll()
		if err != nil {
			panic(err)
		}

		for _, dLine := range dLines {
			fmt.Println(dLine)
			devList = append(devList, dLine[1])
		}

		testerscsv := "data/testers.csv"

		f2, err := os.Open(testerscsv)
		if err != nil {
			panic(err)
		}
		defer f2.Close()

		tLines, err := csv.NewReader(f2).ReadAll()
		if err != nil {
			panic(err)
		}
		skip := true
		for _, tLine := range tLines {
			if skip {
				skip = false
				continue
			}
			var devices []string
			tdevicescsv := "data/tester_device.csv"

			f3, err := os.Open(tdevicescsv)
			if err != nil {
				panic(err)
			}
			defer f3.Close()

			tdlines, err := csv.NewReader(f3).ReadAll()
			if err != nil {
				panic(err)
			}

			for _, tdLine := range tdlines {
				if tdLine[0] == tLine[0] {
					i, _ := strconv.Atoi(tdLine[1])
					devices = append(devices, devList[i])
				}
			}

			id, _ := strconv.Atoi(tLine[0])

			data := m.User{
				ID:        id,
				FirstName: tLine[1],
				LastName:  tLine[2],
				Country:   tLine[3],
				LastLogin: tLine[4],
				Devices:   devices,
			}
			fmt.Println(tLine)
			fmt.Println(data)
			rest.service.CreateUser(data)
		}
	}
	return nil
}

func (rest restService) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	body, err := readRequestBody(r)
	if err != nil {
		jsonerr.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	a := m.User{}
	err = json.Unmarshal(body, &a)

	a, err = rest.service.CreateUser(a)
	if err != nil {
		jsonerr.JSONError(w, err, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(a)
}

func (rest restService) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var users []m.User
	users, err := rest.service.GetUsers(r.URL.Query())
	if err != nil {
		jsonerr.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func (rest restService) GetUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)
	id := vars["id"]

	var a m.User
	a, err := rest.service.GetUserByID(id)
	if err != nil {
		jsonerr.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(a)
}

func (rest restService) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)
	id := vars["id"]

	body, err := readRequestBody(r)
	if err != nil {
		jsonerr.JSONError(w, err, http.StatusInternalServerError)
		return
	}
	var a m.User
	err = json.Unmarshal(body, &a)
	if err != nil {
		jsonerr.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	a, err = rest.service.UpdateUser(id, a)
	if err != nil {
		jsonerr.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(a)
}

func (rest restService) ReplaceUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)
	id := vars["id"]

	body, err := readRequestBody(r)
	if err != nil {
		jsonerr.JSONError(w, err, http.StatusInternalServerError)
		return
	}
	var a m.User
	err = json.Unmarshal(body, &a)
	if err != nil {
		jsonerr.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	a, err = rest.service.ReplaceUser(id, a)
	if err != nil {
		jsonerr.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(a)
}

func (rest restService) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)
	id := vars["id"]

	err := rest.service.DeleteUser(id)
	if err != nil {
		jsonerr.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func readRequestBody(r *http.Request) ([]byte, error) {
	if r.Body == nil {
		return nil, errors.New("No input received ")
	}
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		return body, err
	}
	if err := r.Body.Close(); err != nil {
		return body, err
	}
	return body, err
}
