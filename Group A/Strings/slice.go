package main

import "fmt"

func main() {
	// x:= type{values} composite type
	x := []int{1, 2, 3, 4, 5, 6, 7} //static allocation ie len == cap
	x = append(x, 10, 20)           // this doubles the cap
	fmt.Println(len(x), cap(x))

	y := []int{-1, -2, -3}

	x = append(x, y...)

	fmt.Println(x)
	// for i, v := range x {
	// 	fmt.Println(i, v)
	// }

	//Deleting 4 from slice.
	x = append(x[:3], x[4:]...)
	fmt.Println(x)

	//make slice
	age := make([]int, 4, 10)
	age[3] = 22
	fmt.Println(age)

	age[4] = 244
	fmt.Println(age)



}
