package main

import "fmt"

func totalSum(numbers ...int) int { // need to use any then use interface{}
	total := 0
	for _, num := range numbers {
		total = total + num
	}
	return total
}

func variadicFunction() {
	fmt.Println(2, 3, 2, "hello") // println is a variadic Function, it accept all type of value and infinit time

	nums := []int{1, 3, 4, 5, 2}
	fmt.Println(nums)

	// create custom variadic function totalSum
	value := totalSum(nums...)
	fmt.Println(value)

	//	...int in Function Definition: This allows the function to accept any number of int arguments, treating them as a slice inside the function.
	// nums... in Function Calls: This unpacks a slice into individual arguments, enabling you to pass a slice to a variadic function that expects individual elements.
}
