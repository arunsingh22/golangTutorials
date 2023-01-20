package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/panasonic/golang/middleware"
)

// https://pkg.go.dev/github.com/gorilla/mux#section-readme

// 1. Routes are tested in the order they were added to the router. If two routes match, the first one wins:
// 2. Here, our log function returns a new handler (remember that http.HandlerFunc is a valid http.Handler too)
// that will print the “Before”

func greet(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req) // Key:value pair => map
	name := vars["name"]
	key := vars["key"]
	if key != "" {
		fmt.Println("FOr this middleware logging is working")
	}
	// w.Header().Set("Content-Type", "application/text")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome to the Page %s", name)
}

func main() {
	fmt.Println("Starting server...")
	defer fmt.Println("Stopping server...")
	mx := mux.NewRouter()
	// mx.Use() is used to apply the middle to all the routes
	mx.Use(middleware.GlobalLogging)

	// This way of adding the routes
	mx.HandleFunc("/{name}", middleware.Logging(greet)).Methods("GET")

	http.NewRequest(http.MethodGet,nil,nil)

	// There's one more thing about subroutes. When a subrouter has a path prefix,
	// the inner routes use it as base for their paths:
	// s := mx.PathPrefix("/products").Subrouter()
	// // "/products/"
	// s.HandleFunc("/", greet)
	// // "/products/{key}/"
	// s.HandleFunc("/{key}/", greet)
	// // "/products/{key}/details"
	// s.HandleFunc("/{key}/details", greet)

	http.ListenAndServe(":8080", mx)

}
