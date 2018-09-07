package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct{ width, height int }

func (Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.width, img.height)
}

func (Image) At(x, y int) color.Color {
	c := uint8(x ^ y)
	return color.RGBA{c, c, 255, 255}
}

func main() {
	m := Image{100, 100}
	pic.ShowImage(m)
}
