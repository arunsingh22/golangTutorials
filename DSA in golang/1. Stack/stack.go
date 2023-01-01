package main

import (
	"errors"
	"fmt"
)

type Stack []interface{}

// Isempty checks if stack is empty or not
func (s *Stack) Isempty() bool {
	return len(*s) == 0
}

func (s *Stack) push(x string) error {
	*s = append(*s, x) //appending directly to the stack
	return nil
}

func (s *Stack) pop() error {
	if s.Isempty() == true {
		return errors.New("Stack underflow")
	} else {
		idx := len(*s) - 1
		element := (*s)[idx]
		fmt.Sprintf("Element poped from Stack %v", element)
		*s = (*s)[:idx] // the last element is ALWAYS excluded
		return nil
	}
}

func (s *Stack) display() {
	for i := 0; i < len(*s); i++ {
		fmt.Print((*s)[i], "->")
	}
	fmt.Println()
}

func main() {
	var stk Stack
	stk.push("1")
	stk.push("2")
	stk.display()
	stk.pop()
	stk.display()
}
