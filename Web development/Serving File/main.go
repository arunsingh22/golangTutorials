package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func anotherServer(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Content-type", "text/html;charset= utf-8")
	io.WriteString(w, `
	<!-- Not serving from our Server -->
	<img src ="https://en.wikipedia.org/wiki/Image#/media/File:Image_created_with_a_mobile_phone.png">
	`)
}

func myserver(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("content-type", "text/html;charset= utf-8")
	if f, err := os.Open("linux.jpg"); err == nil {
		defer f.Close()
		io.Copy(w, f)
	} else {
		fmt.Println("ERROR")
	}

}

func main() {
	http.HandleFunc("/anotherServer", anotherServer)
	http.HandleFunc("/myserver", myserver)

	fmt.Println("Server Listen on.. 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
