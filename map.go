package main

import (
	"fmt"
	"maps"
)

func mapFunction() {

	// map -> hash, object, dict on other language

	// create map function using make function
	m := make(map[string]string) // map[key_type]value_type

	m["name"] = "shriyansh"
	m["del"] = "9393"
	fmt.Println(m)
	fmt.Println(len(m))
	fmt.Println(m["name"])
	fmt.Println(m["age"]) // output = ""
	// Note -> if key does not exists in the map then it return zero value

	// ===============================
	delete(m, "del") // delete particular key value from Map
	fmt.Println(m)
	clear(m) // clear all key value from Map
	fmt.Println(m)
	// ===============================

	newMapped := make(map[string]int) // map[key_type]value_type
	newMapped["age"] = 30
	fmt.Println(newMapped)
	fmt.Println(newMapped["name"]) // output = 0
	fmt.Println(newMapped["age"])

	// create map function with out make function
	// check map length and is key exists check

	mapWithoutMake := map[string]int{"price": 2, "amount": 303}
	mapWithoutMake2 := map[string]int{"price": 212, "amount": 211}

	fmt.Println(mapWithoutMake)
	fmt.Println(mapWithoutMake2)

	fmt.Println(maps.Equal(mapWithoutMake, mapWithoutMake2)) // Maps function to check two map key and value both

	value, ok := mapWithoutMake["amount"] // value from the Key, true is key exists in map

	fmt.Println(ok)
	fmt.Println(value)

	if ok {
		fmt.Println("all Ok") // output = 0
	} else {
		fmt.Println("not Ok") // output = 0
	}

}
