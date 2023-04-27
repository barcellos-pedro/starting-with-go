package main

import "fmt"

type Person struct{}

func main() {
	// var inter interface{} = Person{}
	// var inter interface{} = 0.2
	// var inter interface{} = 4
	var inter interface{} = ""

	// switch val := inter.(type) // to store the value
	switch inter.(type) {
	case int:
		fmt.Printf("Type is int: %T\n", inter)
	case string:
		fmt.Printf("Type is string: %T\n", inter)
	case float64:
		fmt.Printf("Type is float: %T\n", inter)
	case Person:
		fmt.Printf("Type is Person: %T\n", inter)
	default:
		fmt.Println("I don't know the type")
	}
}
