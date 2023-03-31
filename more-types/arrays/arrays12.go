package main

import "fmt"

func main() {
	numbers := [3]int{1, 2, 3}
	new_numbers := numbers[:] // new_numbers == numbers

	fmt.Println(numbers)
	fmt.Println(new_numbers)
	new_numbers = append(new_numbers, 4, 5)
	fmt.Println(new_numbers)
}
