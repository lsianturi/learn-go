package main

import (
	"fmt"
	"math"
)

type Vertexy struct {
	X, Y float64
}

func (v Vertexy) Abso() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func (v *Vertexy) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

func main() {
	v := Vertexy{3, 4}
	fmt.Println(v.Abso())
	v.Scale(10)
	fmt.Println(v.Abso())
	v.Scale(2)
	fmt.Println(v.Abso())
}
