package main

import (
	"dev/app/handler"
	"fmt"
	"net/http"
)

func main() {
	router := handler.NewRouter()
	err := http.ListenAndServe(":558", router)
	if err != nil {
		fmt.Println(err)
	}
}
