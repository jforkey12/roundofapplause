package main

import (
	"dev/users/handler"
	"fmt"
	"net/http"
)

func main() {
	router := handler.NewRouter()
	err := http.ListenAndServe(":556", router)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Users service running... RESTAPI port: 556")
}
