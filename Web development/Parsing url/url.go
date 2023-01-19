package main

import (
	"fmt"
	"net/url"
)

const URL = "https://lco.dev:3000/learn?courseName=golang&paymentID=142"

func main() {
	result, err := url.Parse(URL)
	if err != nil {
		panic(err)
	}

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	fmt.Println(result.RawQuery)
	qparms := result.Query() // returns query params in hashtable format

	for key, val := range qparms {
		fmt.Println(key, val)
	}

	// Constructing a URL
	partsOfURL := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}
	fmt.Println(partsOfURL.String())

}
