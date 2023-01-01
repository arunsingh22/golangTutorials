package main

import (
	"fmt"
	"reflect"
)

func main() {
	str := "arun singh"
	fmt.Printf("str's type %T:\n", str)

	for _, c := range str {
		if c == 'a' {
			fmt.Println("Matched")
			fmt.Println("a size:", reflect.TypeOf(c).Size())
		}
		// fmt.Printf("c's value %v type: %T\n", c, c)
		fmt.Println("c size:", reflect.TypeOf(c).Size())

	}

	x := 'a'
	fmt.Println("x type:", reflect.TypeOf(x))
	fmt.Println("x size:", reflect.TypeOf(x).Size())
	fmt.Printf("c's value %v type: %T\n", x, x)

}
