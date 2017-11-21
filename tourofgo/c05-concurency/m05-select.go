package main

import "fmt"

// Fibonacci bla
func Fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("Quit")
			return
		}
	}
}

// main bla
func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 20; i++ {
			fmt.Println(<-c)
		}
		quit <- 999
	}()
	Fibonacci(c, quit)
	fmt.Println("done")
}
