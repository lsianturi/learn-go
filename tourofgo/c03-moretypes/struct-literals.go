package main

import "fmt"

type Vertex1 struct {
	X, Y int
}
type Vertex2 struct {
	X, Y int
}

var (
	v1 = Vertex2{1, 2}
	v2 = Vertex2{X: 1}
	v3 = Vertex2{}
	p  = &Vertex2{1, 2}
)

func main() {
	fmt.Println(v1, v2, v3, p)
	p.X = 5
	fmt.Println(p)
	p := Vertex1{6, 7}
	fmt.Println(p)
}
