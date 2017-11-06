package main

import (
	// "fmt"

	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	a := make([][]uint8, dy)
	// fmt.Println(a)
	for i := range a {
		a[i] = make([]uint8, dx)
	}
	// fmt.Println(a)

	for i := range a {
		for j := range a[i] {
			//a[i][j] = uint8((i * j) / 2)
			//a[i][j] = uint8(i) * uint8(j)
			a[i][j] = uint8(i) ^ uint8(j)
		}
	}
	return a
}

func main() {
	pic.Show(Pic)
}
