package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (ens ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(ens))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := float64(3)
	n := 0
	for {
		z -= (z*z - x) / (2*z)
		fmt.Printf("z is %f\n", z)
		if math.Abs(z*z-x) < 0.00000000000001 {
			break
		}
		n++
	}
	fmt.Printf("count: %d\n", n)
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}


