package main

import "fmt"

func array() {

	// zeroed Value means empty aaray have default values
	// int -> 0, string -> "", bool -> false
	// check the length of array use len() func
	// why we use array -
	// 1- fixed sizes, that is predectable
	// 2- memory optimazation
	// 3- contant time acces

	var numbers [4]int
	fmt.Println(len((numbers)))
	numbers[0] = 5
	fmt.Println(numbers)

	// string have ""
	var names [4]string
	fmt.Println(names)
	names[0] = "shriyansh"
	fmt.Println(names)

	// boolean have false
	var isData [4]bool
	fmt.Println(isData)
	isData[0] = true
	fmt.Println(isData)

	// ======================
	// to declare it in single line values in array

	num := [4]int{2, 4, 5, 2}
	fmt.Println(num)

	// 2D arrays
	number := [2][2]int{{3, 4}, {2, 6}}
	fmt.Println(number)

}
