package main

import "fmt"

// Strategy function type: takes two integers and returns an integer
// This is equivalent to interfac in OOP style
type Operation func(int, int) int

// Context: Holds the currently selected strategy
// Struct which implement the interface
type Calculator struct {
	operation Operation
}

// SetStrategy: Allows changing the active strategy at runtime
func (c *Calculator) SetStrategy(op Operation) {
	c.operation = op
}

// Execute: Uses the current strategy to perform the operation
func (c *Calculator) Execute(a, b int) int {
	if c.operation == nil {
		fmt.Println("No operation set!")
		return 0 // Or handle error appropriately
	}
	return c.operation(a, b)
}

// Concrete Strategy Functions
func add(a, b int) int {
	fmt.Println("Performing addition...")
	return a + b
}

func subtract(a, b int) int {
	fmt.Println("Performing subtraction...")
	return a - b
}

func multiply(a, b int) int {
	fmt.Println("Performing multiplication...")
	return a * b
}

func main() {
	calculator := Calculator{}

	// Set and execute the addition strategy
	calculator.SetStrategy(add)
	result := calculator.Execute(5, 3)
	fmt.Println("Result:", result) // Output: Performing addition... Result: 8

	// Change the strategy to subtraction
	calculator.SetStrategy(subtract)
	result = calculator.Execute(5, 3)
	fmt.Println("Result:", result) // Output: Performing subtraction... Result: 2

	// Change the strategy to multiplication
	calculator.SetStrategy(multiply)
	result = calculator.Execute(5, 3)
	fmt.Println("Result:", result) // Output: Performing multiplication... Result: 15
}
