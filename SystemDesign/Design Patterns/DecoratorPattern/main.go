package main

import "fmt"

// Component Interface
type Coffee interface {
	Cost() int
	Description() string
}

// Concrete Component: Simple Coffee
type SimpleCoffee struct{}

func (s *SimpleCoffee) Cost() int {
	return 5
}

func (s *SimpleCoffee) Description() string {
	return "Simple Coffee"
}

// Decorator Interface
// type CoffeeDecorator interface {
// 	Cost() int
// 	Description() string
// }

// Base Decorator: Wraps a Coffee component
type BeverageDecorator struct {
	coffee Coffee
}

func (b *BeverageDecorator) Cost() int {
	return b.coffee.Cost()
}

func (b *BeverageDecorator) Description() string {
	return b.coffee.Description()
}

// Concrete Decorator: Milk
type MilkDecorator struct {
	BeverageDecorator
}

func (m *MilkDecorator) Cost() int {
	return m.BeverageDecorator.Cost() + 2
}

func (m *MilkDecorator) Description() string {
	return m.BeverageDecorator.Description() + ", Milk"
}

// Concrete Decorator: Sugar
type SugarDecorator struct {
	BeverageDecorator
}

func (s *SugarDecorator) Cost() int {
	return s.BeverageDecorator.Cost() + 1
}

func (s *SugarDecorator) Description() string {
	return s.BeverageDecorator.Description() + ", Sugar"
}

func main() {
	// Create a simple coffee
	myCoffee := &SimpleCoffee{}
	fmt.Println(myCoffee.Description(), "Cost:", myCoffee.Cost())
	// Output: Simple Coffee Cost: 5

	// Decorate with milk
	myCoffeeWithMilk := &MilkDecorator{BeverageDecorator{coffee: myCoffee}}
	fmt.Println(myCoffeeWithMilk.Description(), "Cost:", myCoffeeWithMilk.Cost()) // Output: Simple Coffee, Milk Cost: 7

	// Decorate with sugar
	myCoffeeWithMilkAndSugar := &SugarDecorator{BeverageDecorator{coffee: myCoffeeWithMilk}}
	fmt.Println(myCoffeeWithMilkAndSugar.Description(), "Cost:", myCoffeeWithMilkAndSugar.Cost()) // Output: Simple Coffee, Milk, Sugar Cost: 8

	// Decorate a simple coffee directly with sugar
	myCoffeeWithSugar := &SugarDecorator{BeverageDecorator{coffee: &SimpleCoffee{}}}
	fmt.Println(myCoffeeWithSugar.Description(), "Cost:", myCoffeeWithSugar.Cost()) // Output: Simple Coffee, Sugar Cost: 6

	// You can even stack multiple decorators of the same type
	doubleMilkCoffee := &MilkDecorator{BeverageDecorator{coffee: &MilkDecorator{BeverageDecorator{coffee: &SimpleCoffee{}}}}}
	fmt.Println(doubleMilkCoffee.Description(), "Cost:", doubleMilkCoffee.Cost()) // Output: Simple Coffee, Milk, Milk Cost: 9
}
