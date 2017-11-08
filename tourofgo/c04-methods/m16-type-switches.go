package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("Type unknown for %T\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(21.2)
	do(true)
}