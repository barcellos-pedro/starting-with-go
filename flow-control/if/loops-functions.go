package main

import (
	"fmt"
)

/*
z² − x above is how far away z² is from where it needs to be (x),
and the division by 2z is the derivative of z²,
to scale how much we adjust z by how quickly z² is changing.
This general approach is called Newton's method.
It works well for many functions but especially well for square root
*/
func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(4))
	fmt.Println("------")
	fmt.Println(Sqrt(16))
}
