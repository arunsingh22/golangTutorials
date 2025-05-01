package main

import (
	"fmt"
)

type stack struct {
	stk    []int
	minStk []int
	size   int
}

func InitStack(len int) *stack {
	fmt.Println("init a new stack")
	return &stack{
		stk:    []int{},
		minStk: []int{},
		size:   len,
	}
}

func (s *stack) push(val ...int) {
	s.stk = append(s.stk, val...)

	for _, v := range val {
		if len(s.minStk) == 0 {
			s.minStk = append(s.minStk, v)

		} else if v <= s.minStk[len(s.minStk)-1] {
			s.minStk = append(s.minStk, v)
		}
	}
	fmt.Println("\n After Insert MinStack: ", s.minStk)
}

func (s *stack) pop() int {
	var val int
	if !s.isEmpty() {
		val = s.stk[len(s.stk)-1]
		s.stk = s.stk[:len(s.stk)-1] // remove the last element
	}
	if val == s.minStk[len(s.minStk)-1] {
		s.minStk = s.minStk[:len(s.minStk)-1]
	}
	fmt.Println("\n After pop MinStack: ", s.minStk)
	return -1
}

func (s *stack) peak() int {
	return s.stk[len(s.stk)-1]
}

func (s *stack) isEmpty() bool {
	return len(s.stk) == 0
}

func (s *stack) clear() {
	// s.stk = []int{} // one way to clear
	s.stk = s.stk[:0] // another way to clear
}

func (s *stack) display() {
	for _, val := range s.stk {
		fmt.Print(val, " ")
	}
}

func (s *stack) Minimum() int {
	if len(s.minStk) == 0 {
		return -1
	}
	return s.minStk[len(s.minStk)-1]
}

func main() {
	stk := InitStack(10)
	stk.push(1, 2, 3, 4, 5, 6, 7, -1)
	// stk.push(6, 7)
	// stk.display()

	fmt.Println(stk.pop())
	stk.display()
	// stk.clear()
	// fmt.Println("\nafter clearing the stack")
	// stk.display()
}
