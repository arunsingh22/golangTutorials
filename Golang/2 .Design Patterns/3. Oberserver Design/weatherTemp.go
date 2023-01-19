package main

import (
	"fmt"
	"net/http"
)

// type Handler interface {
// 	ServeHTTP(http.ResponseWriter, *http.Request)
// }

type handler int

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world")
}

func main() {
	var hh handler
	http.ListenAndServe(":9090", hh)
}
