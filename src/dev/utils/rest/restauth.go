package rest

import (
	"fmt"
	"net/http"
	"time"
)

func RestAPIEntryPoint(inner http.Handler, name string) http.Handler {
	fmt.Println(inner)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Needed to allow Chrome to access this API
		start := time.Now()
		fmt.Println(start)

		inner.ServeHTTP(w, r)
	})
}
