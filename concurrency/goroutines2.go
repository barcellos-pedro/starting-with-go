package main

import (
	"fmt"
	"time"
)

func hello(msg string) {
	fmt.Println(msg)
}

// In a concurrent program, the main() is always a default goroutine.
// Other goroutines can not execute if the main() is not executing.
func main() {
	go hello("hello")
	go hello("world")
	// So, in order to make sure that all the goroutines are executed before the main function ends, we sleep the process
	time.Sleep(time.Second * 1)
}
