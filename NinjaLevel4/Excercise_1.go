package main

import "fmt"

func main() {
	//composite literal to create array
	// arr := [5]int{1,2,3,4,5}
	// for i,v := range arr{
	// 	fmt.Println(i,v)
	// }
	// fmt.Printf("%T",arr)

	//creating a slice using composite literal

	xi := []int{1, 2, 3, 4, 10, 15, 25}
	a := []int{1, 2, 3, 4, 5}

	//slicing a slice
	fmt.Println(xi[:])
	fmt.Println(xi[:5])

	xi = append(xi, 55, 10, 11)
	fmt.Println(xi)

	xi = append(xi, a...)
	fmt.Println(xi)

}
