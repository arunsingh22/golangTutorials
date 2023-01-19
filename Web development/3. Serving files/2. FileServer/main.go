package main

import (
	"fmt"
	"net/http"
)

func sendpic(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "linux.jpg")
}

func main() {
	mx := http.NewServeMux()
	fmt.Println("Starting server..")
	// attached routes/handlers
	mx.Handle("/favicon", http.NotFoundHandler())
	mx.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	mx.HandleFunc("/pic", sendpic)

	http.ListenAndServe(":8100", mx)
	fmt.Println("Stopping server..")
}
