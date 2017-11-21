package main

import "fmt"

func fibonacci1() func() int {
	past, curr := 0, 1
	return func() int {
		defer func() {
			past += curr
			past, curr = curr, past

		}()
		return past
	}
}

func fibonacci2() func() int {
	g, f := 0, 1
	return func() int {
		f, g = g, f+g
		return f
	}
}

func main() {
	f := fibonacci1()
	f2 := fibonacci2()
	for i := 0; i < 20; i++ {
		fmt.Println(i, f(), f2())
	}
}
