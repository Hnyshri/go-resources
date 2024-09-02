package main

import "fmt"

// interface{} and any are used to represent a type that can hold any value.
// Use interface{} if you're maintaining or working on older codebases.
// Use any in new code to take advantage of the clearer, more modern syntax.

func printSlice[T any](numbers []T) {
	for _, item := range numbers {
		fmt.Println(item)
	}
}

// comparable is a constraint introduced in Go 1.18 with the advent of generics. It restricts the type to those that can be compared
// using the == and != operators. This includes all primitive types (like int, float64, string, etc.), pointers, and other types that
// support equality comparison.

// Use comparable in generic functions or types when you need to ensure that the types can be compared for equality.
func printSlice3[T comparable](numbers []T) {
	for _, item := range numbers {
		fmt.Println(item)
	}
}

func printSlice2[T int | string | bool](numbers []T) { // this is [T int | string | bool] generic
	for _, item := range numbers {
		fmt.Println(item)
	}
}

// syntax for generic using struct
type stack[T any] struct {
	elements []T
}

// introduce after 1.18 version
func genericsFUnction() {
	numbers := []int{1, 3, 5, 3}
	printSlice(numbers)
	printSlice2(numbers)
	printSlice3(numbers)

	myStack := stack[int]{
		elements: []int{1, 34, 4},
	}
	fmt.Println(myStack)

}
