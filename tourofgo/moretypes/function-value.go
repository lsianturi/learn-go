package main

import (
	"fmt"
	"math"
)

func hitung(fungsi func(float64, float64) float64) float64 {
	return fungsi(5, 12)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	//dua statement ini menghasilkan nilai yang sama
	fmt.Println(hypot(5, 12))
	fmt.Println(hitung(hypot))

	//dua statement ini menghasilkan nilai yang sama
	fmt.Println(hitung(math.Pow))
	fmt.Println(math.Pow(5, 12))
}
