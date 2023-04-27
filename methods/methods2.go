package main

import "fmt"

type MyInt int

// We can declare a method on non-struct types, too.
func (value MyInt) toPositive() MyInt {
	if value < 0 {
		return value * -1
	}
	return value
}

func main() {
	// var value MyInt = -23
	value := MyInt(-23)
	fmt.Println(value.toPositive())

	value = 4
	fmt.Println(value.toPositive())
}
