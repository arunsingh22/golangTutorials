package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

func foo(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("my-cookie")

	if err == http.ErrNoCookie {
		// first time visiter
		cookie = &http.Cookie{
			Name:  "my-cookie",
			Value: "0",
		}
	}

	count, _ := strconv.Atoi(cookie.Value)
	count++
	cookie.Value = strconv.Itoa(count)

	http.SetCookie(w, cookie)
	io.WriteString(w, cookie.Value) //writes to the browser*

	fmt.Println("Cookies Set!")
}

func main() {
	// Cookies are domain specific
	fmt.Println("Cookies App running...")
	http.HandleFunc("/", foo)
	// http.HandleFunc("/readCookies", readCookies)
	log.Fatal(http.ListenAndServe(":9000", nil))
}
