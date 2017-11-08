package main

import "fmt"

func main() {
	var a byte
	a = 'A'
	fmt.Println(string(a))

	a = a + 13
	fmt.Println(string(a))
}
