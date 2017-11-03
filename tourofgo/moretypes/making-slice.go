package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSclice("a", a)
	a[0] = 1
	printSclice("a", a)
	b := make([]int, 0, 5)
	printSclice("b", b)

	b = append(b, 1)
	printSclice("b", b)

	b = append(b, 7)
	printSclice("b", b)

	b[0] = 1000
	printSclice("b", b)

	c := b[:2]
	printSclice("c", c)

	d := c[2:5]
	printSclice("d", d)
	d[1] = 19
	printSclice("d", d)
	printSclice("c", c)
	printSclice("b", b)

}
func printSclice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
