package main

import (
	"log"
	"net/http"
)

func main() {
	// http.Handle("/", http.FileServer(http.Dir(".")))
	// http.ListenAndServe(":9000", nil)

	http.Handle("/", http.StripPrefix("/resource/", http.FileServer(http.Dir("."))))
	log.Fatal(http.ListenAndServe(":9000", nil))
}


