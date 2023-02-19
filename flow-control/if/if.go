package main

import "fmt"

const age int = 18

func isAdultBool(age int) bool {
	if age < 18 {
		return false
	}
	return true
}

func isAdultString(age int) string {
	// if with short statement
	// msg are only available until the end of the if scope
	if msg := "adult"; isAdultBool(age) {
		return msg // adult
	} else {
		// fmt.Println(msg) // msg is available here
	}

	// msg is undefined here

	return "minor"
}

func main() {
	fmt.Println(isAdultBool(age))
	fmt.Println(isAdultString(13))
	fmt.Println(isAdultString(25))
}
