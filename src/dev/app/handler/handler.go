package handler

import (
	m "dev/app/models"
	"dev/utils/jsonerr"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
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

		testerscsv := "data/testers.csv"

		f3, err := os.Open(testerscsv)
		if err != nil {
			panic(err)
		}
		defer f3.Close()

		tLines, err := csv.NewReader(f3).ReadAll()
		if err != nil {
			panic(err)
		}
		skip2 := true
		for _, tLine := range tLines {
			if skip2 {
				skip2 = false
				continue
			}
			var devices []string
			tdevicescsv := "data/tester_device.csv"

			f4, err := os.Open(tdevicescsv)
			if err != nil {
				panic(err)
			}
			defer f4.Close()

			tdlines, err := csv.NewReader(f4).ReadAll()
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

		//create a bunch of resources to test scalability
	} else {
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

		for i := 1; i <= 1000000; i++ {
			data := m.Bug{
				ID:        i,
				Device:    devList[rand.Intn(len(devList))],
				CreatedBy: rand.Intn(100000),
			}
			fmt.Println(data)
			rest.service.CreateBug(data)
		}

		countryList := [...]string{"US", "GB", "JP"}

		for i := 1; i <= 100000; i++ {
			data := m.User{
				ID:        i,
				FirstName: RandStringBytes(8),
				LastName:  RandStringBytes(8),
				Country:   countryList[rand.Intn(len(countryList))],
				LastLogin: RandStringBytes(8),
				Devices:   devList,
			}

			fmt.Println(data)
			rest.service.CreateUser(data)
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

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
