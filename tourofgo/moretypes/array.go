package main

import "fmt"

func main() {
	var a [3]string
	a[0] = "Hello"
	a[1] = "world"
	a[2] = "!"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	var b [2][2]string
	b[0][0] = "hello"
	b[1][1] = "world"
	fmt.Println(b)
}
