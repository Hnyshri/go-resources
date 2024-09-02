package main

import "fmt"

// By value pass as agrument
func changeNumber(numberValue int) { // Passing by Value

	numberValue = 5 // Disadvantage - it will create new value of that variable that means take more memory Address
	fmt.Println("changeNumber Print Value ", numberValue)

}

// By reference pass as agrument
func changeNumber2(numberValue *int) { // Passing by Reference (Using Pointers)

	*numberValue = 5 // *numberValue Dereference value that means change the value of the pointer
	// Advantage - it will take the pointer refernce to change the value so no need to create the variable or memory Allocation
	fmt.Println("changeNumber2 Print Value ", numberValue)

}

func changeSliceFunction(sliceNumber *[]int) { // Modifying a Slice Using Pointers

	*sliceNumber = append(*sliceNumber, 10) // Dereference value
	fmt.Println("changeSliceFunction  Pointer Call sliceNumber ", sliceNumber)
}

func pointers() {

	numberValue := 1

	fmt.Println("Memory Address check of numberValue ", &numberValue)

	changeNumber(numberValue)
	fmt.Println("Before Point Call Print Value ", numberValue)

	changeNumber2(&numberValue) // &numberValue to get the memory address of that variable

	fmt.Println("After Pointer Call Print Value ", numberValue)

	sliceNumber := []int{1, 3, 6}
	changeSliceFunction(&sliceNumber)
	fmt.Println("After Pointer Call sliceNumber ", sliceNumber)

}

// 1- Memory Address: &numberValue gets the memory address of numberValue.
// 2- Passing by Value: Calling changeNumber doesn't change the original numberValue.
// 3- Passing by Reference: Calling changeNumber2 with &numberValue changes the original numberValue.
// 4- Modifying a Slice: The slice sliceNumber is modified by changeSliceFunction, reflecting the change in the original slice.

// Summary
// 1- Pointers in Go hold the memory address of a variable.
// 2- Dereferencing (*pointer) accesses or modifies the value stored at that memory address.
// 3- Passing by reference allows functions to modify the original variable or data structure.
