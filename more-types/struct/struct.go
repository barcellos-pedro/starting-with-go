package main

import "fmt"

// A struct is a collection of fields
type Person struct {
	name   string
	age    int
	status bool
}

func main() {
	pedro := Person{"pedro", 25, false}
	fmt.Println(pedro)
	fmt.Println(pedro.name, pedro.status)
	fmt.Println(Person{"Ana", 26, true})

	// using pointers
	a := Person{"a", 0, false}
	var pointer *Person = &a

	fmt.Println(a.age)
	// with pointers we don't need explicit deference
	// (*pointer).age = 25
	pointer.age = 25
	fmt.Println(a.age)
}
