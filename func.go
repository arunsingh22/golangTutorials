package main

import "fmt"

func main(){
	defer foo()
	xi := []int{1,2,10,11,4,5,6} //slice of int using composite literal
	s := sum(xi...) //unfurling of slice
	fmt.Println("The sum is: ",s)

	
}

//variadic func 0->N
func sum(x ...int) int {
	s := 0
	for _,v := range x{
		s += v
	}
	return s
}
// func (r reciver) identifier(parameters) return TYPE(s) {code....}


func foo(){
	fmt.Println("Executed in last")
}