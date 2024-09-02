package main

import "fmt"

// it is use for iterating over data structure
func rangeFunction() {

	// for slices
	numbers := []int{3, 4, 1, 4}

	// classic for Loop
	for i := 0; i < len(numbers); i++ {
		fmt.Println(i)
		fmt.Println(numbers[i])
	}

	// sum := 0
	//using range
	for index, num := range numbers {
		fmt.Println(index, num)
	}

	// for map
	mapNumbers := map[string]string{"fName": "shriyansh", "lname": "Gupta"}
	for key, value := range mapNumbers {
		fmt.Println(key, value) // print key and value
	}
	for key := range mapNumbers {
		fmt.Println(key) // print key only
	}

	// range in String value
	// unicode -> unicode code point rune
	// key -> starting byte of rune
	for key, unicode := range "shriyansh" {
		fmt.Println(key, unicode, string(unicode)) // print key and value of char in number
	}
}
