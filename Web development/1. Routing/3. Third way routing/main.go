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
	//HandleFunc attaches handlers to defautServeMux
	// HandleFunc is using Adapter pattern 
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/cat", cat)
	fmt.Println("Starting server...at 8080")
	http.ListenAndServe(":8080", nil)
}
