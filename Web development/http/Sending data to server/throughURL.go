package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func foo(w http.ResponseWriter, r *http.Request) {
	v := r.FormValue("q")
	w.Header().Set("Content-type", "text/html;charset=utf-8")
	io.WriteString(w, `
		<form method="get">
		<input type="text name="q">
		<input type="submit">
		</form>
		<br>`+v)
	fmt.Println("Form Value: ", v)
}

func main() {
	fmt.Println("Sending data thrugh URL")
	http.HandleFunc("/", foo)
	log.Fatal(http.ListenAndServe(":9000", nil))
}
