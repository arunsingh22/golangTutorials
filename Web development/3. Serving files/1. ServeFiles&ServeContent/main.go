package main

import (
	"fmt"
	"net/http"
)

func root(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "You are in root dir")
}

func sendpic(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "linux.jpg")
}

func main() {
	// created a new NewServeMux
	mx := http.NewServeMux()
	// attached routes/handlers
	mx.HandleFunc("/", root)
	mx.HandleFunc("/pic", sendpic)

	http.ListenAndServe(":8000", mx)

	// srv := &http.Server{
	// 	Handler: mx, // This is misleading term it's actually expecting a MUX
	// 	Addr:    "127.0.0.1:8000",
	// 	// Good practice: enforce timeouts for servers you create!
	// 	WriteTimeout: 5 * time.Second,
	// 	ReadTimeout:  5 * time.Second,
	// }

	// // blocking stmt so ideally we should run our server in goroutine to unblock
	// // log.Fatal(srv.ListenAndServe())
	// go func() {
	// 	if err := srv.ListenAndServe(); err != nil {
	// 		log.Println(err)
	// 	}
	// }()

}
