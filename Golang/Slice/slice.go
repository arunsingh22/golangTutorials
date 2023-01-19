package main

import "fmt"

// Note: Go doesn't have negative indexing like Python does. This is a deliberate design decision — keeping the language simple can help save you from subtle bugs. Use the index len(a)-1 to access the last element of a slice or array a.

// a := []string{"A", "B", "C"}
// s := a[len(a)-1] // C

// https://yourbasic.org/golang/last-item-in-slice/

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

	age = append(age, 244) // Recommnded way of adding anyting to the slice.
	// age[4] = 244
	// Directly accessing the index beyond the len throws runtime error so always use append to add new elements it automatically takes care of resizing the slice.
	fmt.Println(age)

	fmt.Println("-----Example-1 for copy()-----")
	//https://golangbyexample.com/copy-function-in-golang/
	// copy(dst,src) it takes 2 slices of same type , and copys from src to dst
	// NOTE: the copy is always from 0th index and if there are elements already present in dst slice, they will be replaced by the src elements.
	// NOTE 2 : copy method returns # of elements copied.
	src := []int{1, 2, 3}
	dst := []int{4, 5, 6, 7}
	element_copied := copy(dst, src)
	fmt.Println(element_copied, dst)
}
