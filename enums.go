package main

import "fmt"

// Go doesn't have a built-in enum type like some other programming languages (e.g., C++ or Java). However,
// you can achieve similar functionality using constants and the iota identifier.

// enumrated types

// Define an enumeration using iota
type orderStatus int // create own struct type as a integer

type orderStatusInString string // create own struct type as a String

const (
	Pending orderStatus = iota
	Confirmed
	Delivered
	Cancelled
)

const (
	PENDING   orderStatusInString = "Pending"
	CONFIRMED                     = "Confirmed"
	DELIVERED                     = "Delivered"
	CANCELLED                     = "Cancelled"
)

// iota => is a predeclared identifier representing the untyped integer ordinal number of the current const specification
// in a (usually parenthesized) const declaration. It is zero-indexed.

func changingOrderStatus(status orderStatus) {
	fmt.Println("Order status has been changed to ", status)
}

func changingOrderStatusInString(status orderStatusInString) {
	fmt.Println("Order status has been changed changingOrderStatusInString to ", status)
}

func enumFunction() {
	changingOrderStatus(Pending)
	changingOrderStatusInString(DELIVERED)
}
