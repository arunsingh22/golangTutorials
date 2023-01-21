package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

// name of my handlerFunc is cat
func cat(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Third response from cat")
}

func main() {
	var wg sync.WaitGroup
	// mx := http.NewServeMux()
	// mx := mux.NewRouter()
	http.HandleFunc("/", cat)

	// Creating server with explicit configurations
	// The server takes in a explict mux wheather a NewServeMux() or gorillaMux or defaultServeMux() doesn't matter
	srv := &http.Server{
		Handler: nil, // This is misleading term it's actually expecting a MUX
		Addr:    ":11000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	// blocking stmt so ideally we should run our server in goroutine to unblock
	// log.Fatal(srv.ListenAndServe())
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		} else {
			fmt.Println("Started server..")
		}
	}(&wg)

	wg.Wait()
	fmt.Println("Stopped")

}
