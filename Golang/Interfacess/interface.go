package main

import "fmt"

type area interface {
	Area(int) (int, error)
}
type volume interface {
	volume(int) (int, error)
}

type shape struct {
	a int
}

// Value receiver
func (obj shape) Area() (int, error) {
	return obj.a * obj.a, nil
}

func (obj shape) volume() (int, error) {
	return obj.a * obj.a * obj.a, nil
}

func main() {
	var s shape // struct obj
	var a area  // interface object a
	s = a
	s.a = 2 //violating encapsulation principle

	fmt.Println(obj.Area())

	// No compile-time error BUT Runtime error
	// panic: runtime error: invalid memory address or nil pointer dereference
	// var x area
	// fmt.Println(x.Area(5))

}
