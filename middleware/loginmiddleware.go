package middleware

import (
	"fmt"
	"net/http"
)

func LoginMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Url: ", r.URL)
		h.ServeHTTP(w, r)
	})
}
