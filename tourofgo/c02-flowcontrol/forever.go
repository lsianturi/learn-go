package main

import "fmt"

func main() {
	i := 0
	for {
		fmt.Println("forever")
		i++
		if i == 10 {
			break
		}
	}
}
