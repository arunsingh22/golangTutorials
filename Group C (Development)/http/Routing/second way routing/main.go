package main

import (
	"fmt"
	"net/http"
)

type dog int

func (d dog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from dog")
}

type cat int

func (c cat) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from cat")
}

func main() {
	var d dog
	var c cat

	// No need to create a NewServeMux 
	http.Handle("/dog/", d)
	http.Handle("/cat", c)
	fmt.Println("Starting server...at 8080")
	http.ListenAndServe(":8080", nil) //nil calls defautServeMux
}
