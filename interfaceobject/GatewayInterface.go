package main

import "fmt"

type Gateway interface {
	payment(amount float32)
}

type PhonePay struct{}
type GooglePay struct{}

func (pay PhonePay) payment(amount float32) {
	fmt.Println("We are sending payment by using PhonePay..", amount)
}

func (pay GooglePay) payment(amount float32) {
	fmt.Println("We are sending payment by using GooglePay..", amount)
}

func main() {

	// data preparation for payment gateway

	phonepay := PhonePay{}
	phonepay.payment(4000)

	googlepay := GooglePay{}
	googlepay.payment(3000)

	// For unit testing purpose

	fakepay := FakePay{}
	fakepay.payment(50000)

}

// To create a fake structure data for unit testing

type FakePay struct{}

func (pay FakePay) payment(amount float32) {
	fmt.Println("We are sending payment by using fakePay for unit testing..", amount)
}
