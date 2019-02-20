package handler

import (
	m "dev/bugs/models"
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

		filename := "data/bugs.csv"

		f2, err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		defer f2.Close()

		lines, err := csv.NewReader(f2).ReadAll()
		if err != nil {
			panic(err)
		}
		skip := true
		for _, line := range lines {
			if skip {
				skip = false
				continue
			}
			fmt.Println(line)
			skip = false
			devIndex, _ := strconv.Atoi(line[1])
			id, _ := strconv.Atoi(line[0])
			device := devList[devIndex]
			createdBy, _ := strconv.Atoi(line[2])

			data := m.Bug{
				ID:        id,
				Device:    device,
				CreatedBy: createdBy,
			}
			fmt.Println(data)
			rest.service.CreateBug(data)
		}
	}
	return nil
}

func (rest restService) CreateBug(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	body, err := readRequestBody(r)
	if err != nil {
		jsonerr.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	a := m.Bug{}
	err = json.Unmarshal(body, &a)

	a, err = rest.service.CreateBug(a)
	if err != nil {
		jsonerr.JSONError(w, err, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(a)
}

func (rest restService) GetBugs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var bugs []m.Bug
	bugs, err := rest.service.GetBugs(r.URL.Query())
	if err != nil {
		jsonerr.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bugs)
}

func (rest restService) GetBugByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)
	id := vars["id"]

	var a m.Bug
	a, err := rest.service.GetBugByID(id)
	if err != nil {
		jsonerr.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(a)
}

func (rest restService) UpdateBug(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)
	id := vars["id"]

	body, err := readRequestBody(r)
	if err != nil {
		jsonerr.JSONError(w, err, http.StatusInternalServerError)
		return
	}
	var a m.Bug
	err = json.Unmarshal(body, &a)
	if err != nil {
		jsonerr.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	a, err = rest.service.UpdateBug(id, a)
	if err != nil {
		jsonerr.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(a)
}

func (rest restService) ReplaceBug(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)
	id := vars["id"]

	body, err := readRequestBody(r)
	if err != nil {
		jsonerr.JSONError(w, err, http.StatusInternalServerError)
		return
	}
	var a m.Bug
	err = json.Unmarshal(body, &a)
	if err != nil {
		jsonerr.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	a, err = rest.service.ReplaceBug(id, a)
	if err != nil {
		jsonerr.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(a)
}

func (rest restService) DeleteBug(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)
	id := vars["id"]

	err := rest.service.DeleteBug(id)
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
