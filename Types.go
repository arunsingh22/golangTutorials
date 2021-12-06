package main

import (
	"fmt"
)

// const (
// 	a = 5 //UnTYPED const
// 	b string ="Hi"

// )
func main() {

	// a := 5
	// fmt.Printf("%d ,%b ,%#x\n", a, a, a)
	// b := a << 1
	// fmt.Printf("%d ,%b ,%#x\n", b, b, b)

	// fmt.Println(a,b)

	// c  := `"Raw string literal
	// can span to mulitple lines without
	// any problem"`

	// fmt.Println(c)

	//x := type{values} //composite literal
	x := []int{1, 2, 5, 10,5,11}
	// fmt.Printf("%v %T\n", x, x)
	// fmt.Println(len(x))

	// for i, v := range x {
	// 	fmt.Println(i, v)
	// }



	y := append(x,100,101,102,103)
	fmt.Println(y)

	y  = append(y,x...) //unravales x -> [1, 2, 5, 10,5,11]
	fmt.Println(y)
	

}
