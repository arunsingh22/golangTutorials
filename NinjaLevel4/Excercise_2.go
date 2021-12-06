package main

import "fmt"

func main() {

	xi := make([]string, 50, 50)
	xi = []string{`"arun","singh","how","are","you"`}
	fmt.Println(len(xi))
	fmt.Println(cap(xi))

	//2D slices in Go
	names := []string{"arun", "singh", "Nav"}
	ranks := []string{"Lt.", "Cmd"}
	nameRanks := [][]string{names, ranks}

	for i, v := range nameRanks {
		fmt.Println(i, v)
	}

	//Accessing individual elements with their indexes

}
