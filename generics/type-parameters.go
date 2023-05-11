package main

import "fmt"

// This declaration means that 'list' is a slice
// of any type T that fulfills the built-in constraint comparable
//
// comparable is a useful constraint that makes it possible
// to use the == and != operators on values of the type
func findIndex[T comparable](list []T, value T) int {
	for i, v := range list {
		if v == value {
			return i
		}
	}
	return -1
}

func main() {
	numbers := []int{1, 2, 3}
	num := 4
	fmt.Printf("Index of %v: %v\n", num, findIndex(numbers, num))

	num = 3
	fmt.Printf("Index of %v: %v\n", num, findIndex(numbers, num))

	char := "b"
	letters := []string{"a", "b", "c"}
	fmt.Printf("Index of %v: %v\n", char, findIndex(letters, "b"))
}
