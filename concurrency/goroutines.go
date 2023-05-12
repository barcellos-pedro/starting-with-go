package main

import (
	"fmt"
	"time"
)

func say(msg string) {
	for i := 0; i < 5; i++ {
		time.Sleep(3000 * time.Millisecond)
		fmt.Println(msg)
	}
}

// A Goroutine is a lightweight thread
// managed by the Go runtime
//
// Goroutines run in the same address space
// so access to shared memory must be synchronized
func main() {
	go say("second")
	say("first")
}
