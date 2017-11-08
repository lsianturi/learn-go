package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		if (z * z) == x {
			return z
		}
	}
	return z
}

func main() {
	fmt.Println(sqrt(10))
	fmt.Println(math.Sqrt(10))
}
