package main

import (
	"dev/bugs/handler"
	"fmt"
	"net/http"
)

func main() {
	router := handler.NewRouter()
	err := http.ListenAndServe(":555", router)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Bugs service running... RESTAPI port: 555")
}
