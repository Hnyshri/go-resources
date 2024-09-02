package main

import "fmt"

// Closures in Go are a powerful feature that allows functions to capture and reference variables from their surrounding lexical scope,
// even after the outer function has finished executing. This enables the creation of functions that carry state across multiple invocations.

func counter() func() int {
	count := 0

	return func() int {
		count += 1
		return count
	}
}
func closure() {
	increase := counter() // Creates a new closure with its own count variable
	fmt.Println(increase())
	fmt.Println(increase())
}
