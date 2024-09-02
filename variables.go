package main

import "fmt"

const isUser bool = false

// isNew := "yes" // can not use shorthand syntax out of the funtion

func variables() {

	// Go does not have ternary operator, you will have to use normal if else

	// We can declare a variable inside if construct
	if age := 14; age >= 18 {
		fmt.Println("person is adult")
	} else if age >= 12 {
		fmt.Println("person is not adult")
	}

	var name string = "Shriyansh"

	// infer automatic get the Type
	var name2 = "SHRIYANSH"

	// shorthand syntax to declear the variable
	name3 := "Shriyansh 3 shorthand syntax"

	var name4 string
	name4 = "ETETETE"

	const (
		port = 8081
		host = "localhot"
	)
	fmt.Println(name)
	fmt.Println(name2)
	fmt.Println(name3)
	fmt.Println(name4)
	fmt.Println(port)
	fmt.Println(host)

	var isName bool = true
	fmt.Println(isName)

	var age int = 30
	fmt.Println(age)

	var amount float32 = 50.34222222222222222222
	fmt.Println(amount)

	var amount2 float64 = 21.3333
	fmt.Println(amount2)
	fmt.Println(isUser)

}
