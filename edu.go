package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Sqrt(2.0))

	fmt.Print("Trying Newton's method 10 times: ")
	fmt.Println(Sqrt(2.0, retry10Times))

	fmt.Print("Trying Newton's until a small delta: ")
	fmt.Println(Sqrt(2.0, retryUntilPrecision))
}

// Sqrt computes square root of a number.
func Sqrt(x float64, terminate func(int, float64, float64) bool) float64 {
	var z0, z1 float64
	z0 = x
	z1 = x / 2.0
	for i := 0; i < 1 || !terminate(i, z0, z1); i++ {
		z0 = z1
		z1 = newton(z0, x)
	}
	return z1
}

func retry10Times(i int, z0, z1 float64) bool {
	return i >= 10
}

func retryUntilPrecision(i int, z0, z1 float64) bool {
	return i > 100 || (math.Abs(z0-z1) < .000000001)
}

func newton(zn, x float64) float64 {
	znext := zn - (zn*zn-x)/(float64(2)*zn)
	return znext
}
