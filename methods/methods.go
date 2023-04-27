package main

import "fmt"

type Operations struct {
	a, b int
}

func (ops Operations) add(a, b int) int {
	return a + b
}

func main() {
	// var ops Operations = Operations{}
	ops := Operations{2, 3}
	fmt.Println("value a:", ops.a)
	fmt.Println("value b:", ops.b)

	sum := ops.add(ops.a, ops.b)
	fmt.Println("sum:", sum)

	fmt.Println(ops.add(10, 10))
}
