package main

import "fmt"

func main() {
	names := []string{"pedro", "ana", "ken", "kelly"}

	for index, value := range names {
		fmt.Printf("Index: %d | value: %v", index, value)
		fmt.Println("")
	}

	fmt.Println("")

	// omit index
	for _, value := range names {
		fmt.Printf("value: %v", value)
		fmt.Println("")
	}

	fmt.Println("")

	// omit value
	for index := range names {
		fmt.Println("value:", names[index])
	}
}
