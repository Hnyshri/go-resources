package main

import (
	"fmt"
	"time"
)

type customer struct {
	name   string
	number string
}

type order struct {
	id       int64
	amount   float32
	status   string
	createAt time.Time // nanosecond precision like could not be same
	customer           // struct embadding=> if not use customer then empty Object {}
}

// reciever Type
func (o *order) changeStatus(status string) { // Connect this function to Struct and When changing the value of struct then Use Pointer
	o.status = status // Struct have Dereferencing so no need to declear
}

func (o order) getAmount() float32 { //  When getting the value of struct then not Use Pointer
	return o.amount
}

// create Construct
func newOrders(id int64, amount float32, status string) *order {
	myOrder := order{
		id:     id,
		amount: amount,
		status: status,
	}
	return &myOrder
}

// Customer Data Structure like Class, object used
// In Go, Class have not in Go so use Struct
func structFunction() {

	// if you don't set any field , default value is zero Value
	// int => 0, float => 0, string => "", boolean => false

	// customerData := customer{ // declare the struct Embbading
	// 	name:   "shriyansh",
	// 	number: "9799292992",
	// }

	myOrder := order{ // create the instance of the struct
		id:     2,
		amount: 51.33,
		status: "Completed",
		// customer: customerData,
		customer: customer{ // inline struct embadding function
			name:   "shriyansh",
			number: "9799292992",
		},
	}

	myOrder.createAt = time.Now()

	fmt.Println("Print My Order Data")
	fmt.Println(myOrder)
	fmt.Println(myOrder.amount)

	myOrder.changeStatus("Pending")

	fmt.Println("After changeStatus My Order Data")
	fmt.Println(myOrder)

	fmt.Println("After getAmount My Order Data")
	fmt.Println(myOrder.getAmount())

	myOrder2 := newOrders(5, 43, "Success")
	fmt.Println(myOrder2)

	// inline struct... Only use one time of struct
	language := struct {
		name   string
		isGood bool
	}{"Shriyansh", true}

	fmt.Println("inline struct... Only use one time of struct")

	fmt.Println(language)

}
