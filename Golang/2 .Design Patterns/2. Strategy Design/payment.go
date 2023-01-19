package main

import "fmt"

// PaymentMethod represents a payment method
type PaymentMethod interface {
	Pay(amount float64)
}

// CreditCard represents a credit card payment method
type CreditCard struct{}

// Pay makes a payment using a credit card
func (c *CreditCard) Pay(amount float64) {
	fmt.Printf("Paid %.2f using credit card\n", amount)
}

// Cash represents a cash payment method
type Cash struct{}

// Pay makes a payment using cash
func (c *Cash) Pay(amount float64) {
	fmt.Printf("Paid %.2f using cash\n", amount)
}

// Payment represents a payment
type Payment struct {
	Amount float64
	Method PaymentMethod
}

// MakePayment makes a payment using the specified payment method
func (p *Payment) MakePayment() {
	p.Method.Pay(p.Amount)
}

func main() {
	// Create a credit card payment
	creditCardPayment := Payment{
		Amount: 100,
		Method: &CreditCard{},
	}
	creditCardPayment.MakePayment()

	// Create a cash payment
	cashPayment := Payment{
		Amount: 50,
		Method: &Cash{},
	}
	cashPayment.MakePayment()
}
