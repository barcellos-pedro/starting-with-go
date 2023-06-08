package main

import "fmt"

// not necessary, just to be explicit
type error interface {
	Error() string
}

type MyError struct {
	Cause string
}

// implementing the method
func (err *MyError) Error() string {
	return fmt.Sprintf("MyError error: %v", err.Cause)
}

func sum(a, b int) (int, *MyError) {
	if a == 0 && b == 0 {
		return 0, &MyError{"I won't sum zero values."}
	}

	return a + b, nil
}

func main() {
	if result, err := sum(0, 0); err != nil {
		// fmt.Println(err)
		panic(err)
	} else {
		fmt.Println("Result:", result)
	}
}
