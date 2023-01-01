package main

import (
	"fmt"
	"net/http"
	"net/url"
)

const URL = "https://lco.dev/formdata"

func main() {
	var data url.Values
	data.Add("Name", "Arun singh")
	data.Add("Course", "golang")
	data.Add("Mode", "Online")

	resp, _ := http.PostForm(URL, data)
	fmt.Println(resp)

}
