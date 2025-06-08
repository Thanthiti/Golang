package main

import "fmt"

type Payer interface {
	Pay(amount float64) string
}
type Refunder interface {
	Refund(amount float64) string
}
type CreditCard struct {
	cardNumber string
}

func (c CreditCard) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f via CreditCard %s", amount, c.cardNumber)
}

func (c CreditCard) Refund(amount float64) string {
	return fmt.Sprintf("Refunded %.2f to CreditCard %s", amount, c.cardNumber)
}

type Cash struct{}

func (c Cash) Pay(amount float64) string {
    return fmt.Sprintf("Paid %.2f via Cash", amount)
}

func ProcessPayment(p Payer, amount float64) {
    fmt.Println(p.Pay(amount))
}

func ProcessRefund(r Refunder, amount float64) {
    fmt.Println(r.Refund(amount))
}

func main() {
	cc := CreditCard{cardNumber: "1234-5678-9012-3456"}
    cash := Cash{}

    ProcessPayment(cc, 1000)
    ProcessRefund(cc, 500)
	fmt.Println("------------")
    ProcessPayment(cash, 200)
}