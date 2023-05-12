package main

import (
	"fmt"
	"time"
)

func say(msg string) {
	for i := 0; i < 2; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(msg)
	}
}

// A Goroutine is a lightweight thread
// managed by the Go runtime
//
// Goroutines run in the same address space
// so access to shared memory must be synchronized
func main() {
	go say("world")
	say("hello")
}
