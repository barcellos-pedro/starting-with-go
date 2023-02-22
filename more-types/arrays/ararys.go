package main

import "fmt"

func main() {
	var names [2]string
	names[0] = "Pedro"
	names[1] = "Ana"
	fmt.Println(names)
	fmt.Println(len(names))

	numbers := [3]int{1, 2} // third element will be 0
	fmt.Println(numbers)
}
