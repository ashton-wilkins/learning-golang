package main

import (
	"fmt"
	"math"

	"golang.org/x/tour/pic"
)

func main() {
	exercise02()
}

func exercise02() {
	pic.Show(GraphExtension(10, 10, func(x, y int) uint8 {
		return (uint8(x) + uint8(y)) / uint8(2)
	}))
}

// GraphExtension TODO comment this
func GraphExtension(dx, dy int, fn func(x, y int) uint8) func(int, int) [][]uint8 {
	return func(dx, dy int) [][]uint8 {
		ys := make([][]uint8, dy)
		for y := 0; y < dy; y++ {
			xs := make([]uint8, dx)
			ys[y] = xs
			for x := 0; x < dx; x++ {
				xs = append(xs, fn(x, y))
			}
		}
		return ys
	}
}

func exercise01() {
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
	z1 = z0
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
