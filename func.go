package main

import "fmt"

func main(){
	s := sum(1,2,10,11,4,5,6)
	fmt.Println("The sum is: ",s)
}

//vardiac func
func sum(x ...int) int {
	s := 0
	for _,v := range x{
		s += v
	}
	return s
}
// func (r reciver) identifier(parameters) return TYPE(s) {code....}