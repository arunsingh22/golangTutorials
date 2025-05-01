package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

// ByAge implements sort.Interface based on the Age field.
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// func main() {
// 	family := []Person{
// 		{"Alice", 23},
// 		{"Eve", 2},
// 		{"Bob", 25},
// 	}
// 	sort.Sort(ByAge(family))
// 	fmt.Println(family) // [{Eve 2} {Alice 23} {Bob 25}]
// }

func modify(arr []string) {
	// tmp := make([]string,2)
	fmt.Println("Before appending: ", arr)
	arr = append(arr, "c", "d")
	fmt.Println("After appending: ", arr)
}

func main() {
	// Example 1:
	// slice := []string{"a", "a"}
	// func(slice []string) {
	// 	slice = append(slice, "a")
	// 	slice[0] = "b"
	// 	slice[1] = "b"
	// 	fmt.Print(slice)
	// 	fmt.Println("")
	// }(slice)
	// fmt.Print(slice)

	// Example : 2
	// slice := make([]string, 1, 3)
	// func(slice []string) {
	// 	slice = slice[1:3]
	// 	slice[0] = "b"
	// 	slice[1] = "b"
	// 	fmt.Print(len(slice))
	// 	fmt.Print(slice)
	// }(slice)
	// fmt.Print(len(slice))
	// fmt.Print(slice)

	// slice = []string{"a", "b"}
	// fmt.Println(slice, len(slice), cap(slice))
	// modify(slice)
	// fmt.Println(slice, len(slice), cap(slice))

}
