package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)
	fmt.Printf("%s\n", t)
	// fmt.Println("Location: ", t.Location(), " Time: ", t)
	fmt.Println(t.Format("2006-01-02"))
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Printf("Good evening")
	}
}
