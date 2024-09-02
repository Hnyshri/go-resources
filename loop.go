package main

import "fmt"

// In go only have for Loop

func loop() {

	// make while loop using for loop
	i := 1

	for i <= 3 {
		fmt.Println("i_", i)
		i++
	}

	// ================
	// classic for loop

	for i := 0; i <= 3; i++ {
		if i == 2 {
			continue
		}
		fmt.Println("i_", i)
	}

	// ===================
	// Range in loop introduce in 1.22

	for i := range 10 {
		fmt.Println("i_", i)
	}

	// =================

	// infinite loop in for
	// for {
	// 	fmt.Println("helo")
	// }
}
