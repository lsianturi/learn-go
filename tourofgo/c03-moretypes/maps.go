package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.399678,
	}
	m["Next"] = Vertex{1, 2}
	//fmt.Println(m)
	//fmt.Println(m["Bell Labs"])
	//fmt.Println(m["Next"])

	for i, ver := range m {
		fmt.Printf("%v %v\n", i, ver)
	}
}
