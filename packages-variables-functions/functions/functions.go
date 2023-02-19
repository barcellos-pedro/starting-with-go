package main

import "fmt"

func hello(msg string) string {
	if msg == "" {
		msg = "world"
	}
	return "hello " + msg
}

// we leave only the last type if variables share the same type
func add(x, y int) int {
	return x + y
}

func swap(a, b string) (string, string) {
	return b, a
}

// naked return
func split() (x, y int) {
	x = 1
	y = 2
	return
}

func main() {
	fmt.Println(hello(""))
	fmt.Println(hello("pedro"))
	fmt.Println(add(3, 3))
	word1, word2 := swap("hello", "world")
	fmt.Println(word1, word2)
	fmt.Println(split())
}
