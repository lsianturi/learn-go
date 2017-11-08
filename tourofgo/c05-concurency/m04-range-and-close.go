package main

import "fmt"

func fibonacci(n int, c chan int) {
	past, curr := 0, 1
	for i := 0; i < n; i++ {
		c <- past
		past, curr = curr, curr+past
	}
	close(c)
}

func main() {
	c := make(chan int, 20)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Print(i, " ")
	}
}
