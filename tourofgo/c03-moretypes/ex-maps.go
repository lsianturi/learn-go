package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	// str := strings.Fields(s)
	m := make(map[string]int)

	for _, w := range strings.Fields(s) { //don't need the index
		_, ok := m[w]
		if ok {
			m[w]++
		} else {
			m[w] = 1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
