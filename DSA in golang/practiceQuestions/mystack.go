package main

import "fmt"

type stack []int

func (s *stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *stack) IsFull(top int) bool {
	return len(*s)-1 == top
}

func (s *stack) push(x int) {
	*s = append(*s, x)
}

func (s *stack) pop() {
	if s.IsEmpty() {
		fmt.Println("Underflow")
	}
	top := len(*s) - 1
	fmt.Println("Top element: ", (*s)[top])
	*s = (*s)[:top]
}

func (s *stack) display() {
	for i := 0; i < len(*s); i++ {
		fmt.Print((*s)[i], "->")
	}
	fmt.Println()
}

func main() {
	var stk stack
	stk.push(1)
	stk.push(2)
	stk.push(3)
	stk.push(4)
	stk.display()

	stk.pop()
	stk.push(10)
	stk.display()

}
