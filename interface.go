package main

import "fmt"

// an interface is a type that specifies a set of method signatures. When a type (a struct, for example) implements all the methods
// in an interface, it is said to satisfy the interface. Unlike in some other languages, types do not need to explicitly declare that
// they implement an interface; it happens automatically when the methods are defined.

// for naming to add "er" on last world
type paymenter interface {
	pay(amount float32)
	refund(amount float32, accountNumber string)
}

type payment struct {
	gateway paymenter
}

// open close principle - we are expend the code not modified
func (p payment) makePayment(amount float32) {

	// First Way
	razorPayGateway := razorPay{}
	razorPayGateway.pay(amount)

	// second Way
	stripePayGateway := stripe{}
	stripePayGateway.pay(amount)

	// Third Way - problem to change the  payment Method stripe to razorpay
	p.gateway.pay(amount)
}

type razorPay struct{}

func (r razorPay) pay(amount float32) {
	// logic of the payment or call razorPay Api
	fmt.Println("Make logic of the payment or call razorPay Api", amount)
}

type stripe struct{}

func (s stripe) pay(amount float32) {
	fmt.Println("Make logic of the payment or call stripe Api", amount)
}

func (f stripe) refund(amount float32, accountNumber string) {
	fmt.Println("refund using stripe", amount, accountNumber)
}

type fakePayment struct{}

func (f fakePayment) pay(amount float32) {
	fmt.Println("Make logic of the payment using FakePayment", amount)
}

func interFaceFunction() {

	// newPayment := payment{}
	// newPayment.makePayment(100)

	stripePayGateway2 := stripe{}
	// razorPayPayGateway2 := razorPay{}
	// fakePaymentPayGateway2 := fakePayment{}

	// Problem statement to use multiple payment Gateway on the single time or multiple time so that we use interface for problem solve

	newPayment2 := payment{
		gateway: stripePayGateway2,
	}
	newPayment2.makePayment(100)
}
