package main

import (
	"fmt"
	"slices"
)

// slice -> dynamic arrays, not define the array size
// most used contruct in golang
// + usefull methods
func slicesDefination() {

	// uninitialized slice is nil/null
	var number []int
	fmt.Println(number)
	fmt.Println(len(number))
	fmt.Println(number == nil)

	// ==========================
	// make is pre define func

	// var num = make([]int, 2, 4)
	// var arr = make([]int, 0, 4) // always decleare slice
	num := []int{}

	fmt.Println(len(num))
	fmt.Println(cap(num))   // capacity -> maximun number of element can fit, it will double after the capacity is full
	fmt.Println(num == nil) // make is not nil when define the size in slice only have default value zeroed
	fmt.Println(num)
	num = append(num, 2)
	num = append(num, 5)
	num = append(num, 26)

	fmt.Println(num)
	fmt.Println(cap(num)) // now the cap is double

	// copy slice function
	var arr = make([]int, 0, 4)
	arr = append(arr, 2)
	fmt.Println(arr)

	var arr2 = make([]int, len(arr))
	copy(arr2, arr) // copy slice function
	fmt.Println(arr, arr2)

	// Slice function
	var sliceFunc = []int{3, 4, 2, 122, 3, 4}

	fmt.Println(sliceFunc[1:2]) // Create a sub-slice from index 1 to 2 like array[start:end]
	fmt.Println(sliceFunc[:2])  // Create a sub-slice from index 0 to 2
	fmt.Println(sliceFunc[1:])  // Create a sub-slice from index 1 to end

	// Slices packages

	var SlicesPack = []int{3, 4}
	var SlicesPack2 = []int{3, 4}
	fmt.Println(slices.Equal(SlicesPack, SlicesPack2))

	// 2D array Slices
	var SlicesValues = [][]int{{3, 4}, {2, 1, 2, 3}}
	fmt.Println(SlicesValues)

}
