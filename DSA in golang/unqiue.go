package main

import (
	"fmt"
)

type Employee struct {
	Name   string
	Age    int
	Salary float32
}

// func uniqueTesting() {
// 	emp1 := Employee{"arun", 30, 122.3}
// 	emp2 := Employee{"", 30, 122.3}

// 	h1 := unique.Make(emp1)
// 	h2 := unique.Make(emp2)
// 	if h1 == h2 {
// 		fmt.Println("Both are same!")
// 	} else {
// 		fmt.Println("Both are different!")
// 	}

//		fmt.Println(h1.Value().Age)
//	}
func main() {
	s := make([]int, 0, 2)

	doSomething(s)
	fmt.Println(s[0], s[1])
	fmt.Println(s)
}

func doSomething(a []int) {
	a = append(a, 1)
}
