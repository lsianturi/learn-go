package main

import "fmt"

const Pi = 3.14

func main() {
	const Word = "世界"
	fmt.Println("Hello", Word)
	fmt.Println("Happy", Pi, "Day")

	const Truth = false
	fmt.Println("Go rules?", Truth)
}
