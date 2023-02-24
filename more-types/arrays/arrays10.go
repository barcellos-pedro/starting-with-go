package main

import "fmt"

var names = []string{"pedro", "ana", "ken", "kelly"}

// copy
func main() {
	namesCopy := make([]string, cap(names))
	fmt.Println(len(namesCopy), cap(namesCopy), namesCopy)

	copy(namesCopy, names)

	namesCopy = namesCopy[2:]
	fmt.Println(len(namesCopy), cap(namesCopy), namesCopy)
}
