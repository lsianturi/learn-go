package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	s[0] = 1
	if s == nil {
		fmt.Println("nil!")
	}
}
