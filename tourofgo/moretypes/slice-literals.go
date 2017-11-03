package main

import "fmt"

func main() {
	q := [8]int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)
	fmt.Println("size:", len(q), " capacity:", cap(q))

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{1, true},
		{2, false},
		{3, false},
		{4, true},
		{5, false},
	}
	fmt.Println(s)
}
