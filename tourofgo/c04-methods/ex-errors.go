package main

import (
	"fmt"
	"math"
)

// ErrNegativeSqrt bla
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %g\n", float64(e))
}

// Sqrt bla
func Sqrt(x float64) (float64, error) {
	if x < 1 {
		return 0, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil

}

func main() {
	a, _ := Sqrt(9)
	fmt.Println(a)
	fmt.Println(Sqrt(4))
	fmt.Println(Sqrt(-4))
}
