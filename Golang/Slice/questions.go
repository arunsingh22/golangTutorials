package main

import "fmt"

// Key Observations/Take Aways
// 1. When passing a slice the header gets copied and weather the elements in the underlying
// array will be accessed or not is determined by the len and cap in the sliceHeader
// 2. If you 100% want the slice modification to reflect outside the func ALWAYS PASS BY REF
// 3. IF you pass simply pass the slice, then only when elements which are update/accessed by
// the index value will be reflected out the func.
// 4. In case of a append in the func
// - After appending if it exceed the cap of the original array then 100% a new array is
//   init and the slice header is updated
// If the append func doesn't changes the cap of the original slice it will append into the
// same array and updates it's local sliceHeader and therefore it wont be reflected outside the
// func
// Note Formula:
// y := x[m:n]
// len(y) = m-n
// cap(y) = cap(x)-m

func GuessOuput() {
	a := []int{1, 2, 3}
	b := append(a[:1], 10)
	fmt.Println(b, a)

	// Actual Output
	// b = [1,10]
	// a = [1,10,3]
}

func fullSlicing() {
	a := []int{1, 2, 3}
	fmt.Println(a)
	a = a[:] // a[0:len(a)]
	fmt.Println(a)
	b := a[:0] // this resets the len keeping the cap and underlying array intact
	fmt.Println(b)
	// b[0] = 1000 // since the len(b) 0 , we cannot do this
	b = append(b, 1000) // this works and it changes the underlyign array
	fmt.Println(a)
}

func GuessOutput2() {
	s := make([]int, 0, 5)
	fmt.Printf("s address: %p\n", &s)

	appendWithCap(s)
	fmt.Println(s)
	s = s[:3]
	fmt.Printf("s address: %p\n", &s)
	fmt.Println(s, len(s), cap(s))
}

func appendWithCap(a []int) {
	fmt.Printf("a address: %p\n", &a)
	a = append(a, []int{1, 2, 3}...)
}

func appendBeyondCap(a []int) {
	fmt.Printf("a address: %p\n", &a)
	a = append(a, []int{1, 2, 3, 4, 5, 6}...)
}
