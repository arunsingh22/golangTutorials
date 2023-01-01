package main

import (
	"fmt"
	"net/http"
)

// name of my handlerFunc is dog
func dog(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Third Response from dog")
}

// name of my handlerFunc is cat
func cat(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Third response from cat")
}

func main() {
	//http.HandleFunc and http.Handle both attaches handlers to defautServeMux
	// http.HandlerFunc() this is a type casting of dog into Handler type.
	http.Handle("/dog/", http.HandlerFunc(dog))
	http.Handle("/cat", http.HandlerFunc(cat))
	fmt.Println("Starting server...at 8080")
	http.ListenAndServe(":8080", nil)
}
