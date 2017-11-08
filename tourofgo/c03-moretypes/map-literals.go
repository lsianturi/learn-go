package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell":   Vertex{1, 2},
	"Google": Vertex{3, 4},
}

var n = map[string]Vertex{
	"Bell":   {1, 2},
	"Google": {3, 4},
}

func main() {
	fmt.Println(m)
	fmt.Println(n)
}
