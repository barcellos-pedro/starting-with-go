package main

import "fmt"

// type Interface interface{}
// var inter Interface

// Type assertion provides access to an interface value's underlying concrete value.
func main() {
	var inter interface{}

	value := 2
	inter = value

	// this statement will trigger a panic
	// val := inter.(string)
	// fmt.Println(val)

	if val, ok := inter.(string); ok {
		fmt.Println("Type assertion ok:", val)
	} else {
		fmt.Println("Type assertion false")
	}
}
