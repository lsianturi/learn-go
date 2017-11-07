package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertexo struct {
	X, Y float64
}

func (v *Vertexo) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertexo{3, 4}

	a = f // a MyFloat implements Abser
	fmt.Printf("%+v\n", a)
	fmt.Println(a.Abs())

	a = &v // a *Vertex implements Abser
	fmt.Printf("%+v\n", a)
	fmt.Println(a.Abs())

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	a = v
	fmt.Println(a.Abs())
}
