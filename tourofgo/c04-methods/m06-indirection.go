package main

import "fmt"

type Vertexo struct {
	X, Y float64
}

func (v *Vertexo) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

func ScaleFunc(v *Vertexo, f float64) {
	v.X *= f
	v.Y *= f
}

func main() {
	v := &Vertexo{3, 4}
	fmt.Println(v)
	v.Scale(2)
	fmt.Println(v)
	ScaleFunc(v, 10)
	fmt.Println(v)

	p := Vertexo{4, 3}
	fmt.Println(p)
	p.Scale(3)
	fmt.Println(p)
	ScaleFunc(&p, 10)
	fmt.Println(p)
}
