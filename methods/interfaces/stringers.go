package main

import "fmt"

type Person struct {
	name string
}

// A Stringer is a type that can describe itself as a string.
type Stringer interface {
	String() string
}

// Equivalent to .toString() in other languages
// without this method Println would show {Pedro}
func (person Person) String() string {
	return fmt.Sprintf("Your name is: %v", person.name)
}

func main() {
	p := Person{"Pedro"}
	// fmt.Println(p.String())

	// The fmt package (and others) look for this interface to print values.
	fmt.Println(p)

	// using interface
	var stringer Stringer = p
	fmt.Println(stringer)
}
