package main

import "fmt"

func main() {
	var sum int = 0

	//  for structure
	// init statement, condition expression, post statement
	for i := 0; i < 3; i++ {
		sum += i
	}

	fmt.Println(sum)

	// for continued
	// init and post are optional
	sum = 1

	// for is Go's while
	for sum < 1000 {
		sum += sum
	}

	fmt.Println(sum)

	// infinite loop
	// for {
	// }
}
