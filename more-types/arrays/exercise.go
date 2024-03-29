package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	image := make([][]uint8, dy)

	for i := 0; i < dy; i++ {
		element := make([]uint8, dx)

		for j := 0; j < dx; j++ {
			element[j] = uint8(dx * dy)
		}

		image[i] = element
	}

	return image
}

func main() {
	pic.Show(Pic)
}
