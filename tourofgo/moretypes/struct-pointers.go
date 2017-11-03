package main

import "fmt"

type Vertex1 struct {
	X int
	Y int
}

func main() {
	v := Vertex1{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(*p)
	fmt.Println(v)
}
