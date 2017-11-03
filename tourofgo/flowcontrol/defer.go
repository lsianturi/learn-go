package main

import "fmt"

func deferMulti() {
	fmt.Println("Counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}

func main() {
	defer fmt.Println("world")

	defer deferMulti()
	fmt.Println("Hello")
}
