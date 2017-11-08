package main

import (
	"fmt"
	"math"
)

type Vertexi struct {
	X, Y float64
}

func (v *Vertexi) Scale(f float64) {
	v.X *= f
	v.Y *= f
}
func (v *Vertexi) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertexi{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}
