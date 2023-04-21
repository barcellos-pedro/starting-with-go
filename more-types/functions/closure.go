package main

import "fmt"

func adder(by int) func(int) int {
	return func(value int) int {
		return value + by
	}
}

func main() {
	by10 := adder(10)
	fmt.Println(by10(0))
	fmt.Println(by10(2))

	by100 := adder(100)
	fmt.Println(by100(0))
	fmt.Println(by100(100))
}
