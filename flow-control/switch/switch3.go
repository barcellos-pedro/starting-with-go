package main

import (
	"fmt"
	"time"
)

func test(condition bool) string {
	switch condition {
	case true:
		return "truthy"
	default:
		return "default"
	}
}

func main() {
	t := time.Now()
	// Switch without a condition is the same as switch true.
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	fmt.Println(test(false))
	fmt.Println(test(true))
}
