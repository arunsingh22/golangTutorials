package main

import (
	"fmt"
	"log"
	"net/http"
)

func setCookies(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookie",
		Value: "arun singh",
	})

	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookie-2",
		Value: "shilpa yadav is by bff only when i am not angry ",
	})
	fmt.Println("Cookies Set!")
}
func readCookies(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("my-cookie-2")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Printf(cookie.Value)
	}

	cookie, err = r.Cookie("my-cookie")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Printf(cookie.Value)
	}
}

func main() {
	// Cookies are domain specific
	fmt.Println("Cookies App running...")
	http.HandleFunc("/setCookies", setCookies)
	http.HandleFunc("/readCookies", readCookies)
	log.Fatal(http.ListenAndServe(":9000", nil))
}
