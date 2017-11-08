package main

import (
	"fmt"
	"math"
)

type Vertexy struct {
	X, Y float64
}

func Abs(v Vertexy) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertexy, f float64) {
	v.X *= f
	v.Y *= f
}

func main() {
	v := Vertexy{3, 4}
	fmt.Println(Abs(v))
	Scale(&v, 10)
	fmt.Println(Abs(v))
}
