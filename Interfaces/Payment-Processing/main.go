package main

import "fmt"

type CreditCardProcessor interface {
	ProcessCreditCard(amount float64, cardNumber string)
}

type PayPalProcessor interface {
	ProcessPayPal(amount float64, accountEmail string)
}

type PaymentProcessor interface {
	CreditCardProcessor
	PayPalProcessor
}

type Payment struct{}

func (p Payment) ProcessCreditCard(amount float64, cardNumber string) {
	fmt.Println(amount, cardNumber)
}

func (p Payment) ProcessPayPal(amount float64, accountEmail string) {
	fmt.Println(amount, accountEmail)
}

func PrintPayments(p PaymentProcessor, amount float64, cardNumber string, accountEmail string) {
	p.ProcessCreditCard(amount, cardNumber)
	p.ProcessPayPal(amount, accountEmail)
}

func main() {
	p := Payment{}
	PrintPayments(p, 1200, "3535232", "ankita@gmail.com")
}
