package main

import (
	"fmt"
	"runtime"
)

func main() {
	age := 18
	isAdult := age >= 18

	switch isAdult {
	case true:
		fmt.Println("adult")
	case false:
		fmt.Println("minor")
	}

	// shorthand like if statement
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OX X")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Printf("%s\n", os)
	}

	// fmt.Println(os) // os is undefined here
}
