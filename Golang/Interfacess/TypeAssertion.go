package main

import "fmt"

type cook interface {
	cooking()
	washing()
}

// type switching
func explain(i interface{}) {
	switch v := i.(type) {
	case cook:
		fmt.Println("type is cook", v) //type assertion
	case string:
		fmt.Println(len(v))
	default:
		fmt.Println("No match", v)
	}
}

func main() {
	var c cook
	explain("arun")
	explain(c)
}
