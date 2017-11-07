package main

import (
	"fmt"
	"math"
)

type Vertexa struct {
	X, Y float64
}

func (v Vertexa) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertexa) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertexa{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertexa{3, 4}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}
