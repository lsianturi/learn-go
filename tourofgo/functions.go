package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func tambah(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(5, 6))
	fmt.Println(tambah(11, 44))
}
