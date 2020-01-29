package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {

	slice := make([][]uint8, dx)

	for z := 0; z < dx; z++ {
		slice[z] = make([]uint8, dy)
	}

	for i := 0; i < dx; i++ {
		for z := 0; z < dy; z++ {

			slice[i][z] = uint8((i * z) / 2)
		}
	}

	return slice
}

func main() {
	pic.Show(Pic)
}
