package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const URL = "https://lco.dev"

func main() {
	res, err := http.Get(URL)
	if err != nil {
		panic(err)
	}

	// To list the package dependencies
	// go list -m all : list all the package dependecies
	// go mod graph

	// fmt.Printf("Type of reponse: %T\n  val: %v", res, string(res.Body))
	fmt.Printf("Type of reponse: %T\n", res)
	defer res.Body.Close()

	databyte, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(databyte))
}
