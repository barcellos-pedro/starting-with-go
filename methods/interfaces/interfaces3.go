package main

import "fmt"

type I interface {
	show()
}

// A nil interface value holds neither value nor concrete type.
func main() {
	var i I
	fmt.Println(i == nil)
}
