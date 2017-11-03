package main

import "fmt"

func printSclice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	var s []int
	printSclice(s)
	//append works on nil slices
	s = append(s, 0)
	printSclice(s)

	s = append(s, 1)
	printSclice(s)

	//append more than one at once
	s = append(s, 1, 2, 3, 4, 5)
	printSclice(s)
}
