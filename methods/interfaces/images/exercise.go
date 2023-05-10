package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

// Image unterface to implement its methods on our own Image type
// https://pkg.go.dev/image#Image

// Own image type
type Image struct {
	width, height int
}

func (Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.width, i.height)
}

func (Image) At(x, y int) color.Color {
	R := uint8(100)
	G := uint8(50)
	B := uint8(255)
	A := uint8(255)
	return color.RGBA{R, G, B, A}
}

func main() {
	m := Image{50, 50}
	pic.ShowImage(m)
}
