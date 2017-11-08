package main

import "fmt"

func fibonacci() func() int {
	past, curr := 0, 1
	return func() int {
		defer func() {
			past += curr
			past, curr = curr, past

		}()
		return past
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Println(i, f())
	}
}
