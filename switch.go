package main

import (
	"fmt"
	"time"
)

func switchCase() {

	// Simple switch case
	i := 1
	switch i {
	case 0:
		fmt.Println("yes")
	case 1:
		fmt.Println("yes not")

	default:
		fmt.Println("not")
	}

	// multiple condition switch Case

	fmt.Println(time.Now())
	fmt.Println(time.Now().Weekday())

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's weekend")
	default:
		fmt.Println("it's workday")
	}

	// Type switch
	// go use only double String  - " "
	// whoAmI := func(i any) { //
	whoAmI := func(i interface{}) { // interface{} that mean empty interface accept any value
		switch i.(type) {
		case int:
			fmt.Println("it is integer")
		case string:
			fmt.Println("it is string")
		case bool:
			fmt.Println("it is boolean")
		default:
			fmt.Println("it is other")
		}
	}

	whoAmI(true)
	whoAmI("true")
	whoAmI(22)

}
