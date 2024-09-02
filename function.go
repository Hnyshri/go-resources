package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func sum2(a, b int) int {
	return a + b
}

func getLanguages() (string, bool, int) {
	return "shriyansh", true, 33
}

func processIt(fnc func(a int) int) {
	fmt.Println("ffunc processIt(fn func(a int) int) {")
	fnc(1)
}

func processValue() func(a int) int {
	fmt.Println("func processValue() func(a int) int {")
	return func(a int) int { // it is anonymous function
		fmt.Println("return func(a int) int {")
		return 100
	}
}

func functionDefine() {
	fmt.Println(sum(3, 1))
	fmt.Println(sum2(3, 1))

	val1, val2, val3 := getLanguages()
	fmt.Println(val1, val2, val3)

	value1, _, value3 := getLanguages() // _ use for we are not use that value if decleare
	fmt.Println(value1, value3)

	// -===========================
	fn := func(a int) int {
		fmt.Println("fn := func(a int) int { call")

		return 4
	}
	fmt.Println("processIt")
	processIt(fn)

	// -===========================
	fmt.Println("processValueeeeeeee")
	processValue()
}
