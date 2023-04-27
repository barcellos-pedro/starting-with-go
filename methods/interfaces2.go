package main

import "fmt"

type Person struct {
	name string
}

type I interface {
	show()
}

// Interface values with nil underlying values
func (p *Person) show() {
	if p == nil {
		fmt.Println("receiver is nil")
	} else {
		fmt.Println("receiver not nil", p)
	}
}

func main() {
	var inter I = &Person{}
	inter.show()

	var nilPerson *Person
	inter = nilPerson
	inter.show()

	// an interface value that holds a nil concrete value is itself non-nil.
	fmt.Println("Is interface nil?", inter == nil)
}
