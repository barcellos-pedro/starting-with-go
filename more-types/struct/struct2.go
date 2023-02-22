package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	var v4 Vertex = Vertex{Y: 3}

	fmt.Println(v1, v2, v3, v4)
	fmt.Println(p, *p, p.X)
}
