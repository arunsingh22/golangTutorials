package main

import (
	"fmt"
	"log"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello from root endpoint")
	http.Redirect(w, r, "/redirected", http.StatusMovedPermanently)
}
func redirected(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello from redirected endpoint")
}

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/redirected", redirected)
	log.Fatal(http.ListenAndServe(":9000", nil))
}
