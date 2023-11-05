package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(3)
	n := 0
	for {
		z -= (z*z - x) / (2*z)
		fmt.Printf("z is %f\n", z)
		if math.Abs(z*z-x) < 0.00000000000001 {  // you can choice this number
			break
		}
		n++
	}
	fmt.Printf("count: %d\n", n)
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
