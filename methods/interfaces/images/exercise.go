package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	width, heigth int
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.width, i.heigth)
}

func (i Image) At(x, y int) color.Color {
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
