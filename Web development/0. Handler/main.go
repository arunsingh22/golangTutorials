package main

import (
	"fmt"
	"net/http"
)

// what is handler ?
type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

type arun int

func (obj arun) ServeHTTP(r http.ResponseWriter, req *http.Request) {
	fmt.Println("Hello world")
}
func about(r http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(r, "About page")
}

func main() {
	var obj arun
	// attaching handler to routes
	mx := http.NewServeMux()
	mx.Handle("/", obj)
	mx.HandleFunc("/about", about)
	http.ListenAndServe(":9090", mx) // using defaultServeMux
}
