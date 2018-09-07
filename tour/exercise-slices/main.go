package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	image := make([][]uint8, dy)
	for row := range image {
		image[row] = make([]uint8, dx)
		for col := range image[row] {
			image[row][col] = uint8(row * col)
		}
	}

	return image
}

func main() {
	pic.Show(Pic)
}
