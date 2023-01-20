package middleware

import (
	"fmt"
	"net/http"
)

func Logging(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Local Logging Before")
		next.ServeHTTP(w, r)
		fmt.Println("Local Logging After")
	})
}

func GlobalLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Global Logging Before")
		next.ServeHTTP(w, r)
		fmt.Println("Global Logging After")
	})
}
