package main

import "fmt"

type I interface {
	M() string
}
type T struct {
	S string
}

func (t T) M() string {
	return t.S
}

func main() {
	var i I = T{"Hello"}
	fmt.Println(i.M())
}
