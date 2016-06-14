package main

import (
	"bytes"
	"fmt"
	"math"
	"sync"
	"unicode"

	"golang.org/x/tour/pic"
)

func main() {
	exercise04()
}

func exercise04() {
	f := fibonacci()
	for i := 0; i < 100; i++ {
		fmt.Println(f())
	}
}

func fibonacci() func() int {
	n := 0
	result := 0
	cache := make(map[int]int)
	cache[0] = 0
	cache[1] = 1
	memoize := func() { cache[n] = result; n++ }
	return func() int {
		defer memoize()
		if n <= 0 {
			result = 0
		} else if n == 1 {
			result = 1
		} else {
			result = cache[n-1] + cache[n-2]
		}
		return result
	}
}

func exercise03() {
	wc := WordCount("hello! this is a test to see if the word count is correct or not.")
	fmt.Println(wc)
}

// WordCount calculates the number of words in a string.
func WordCount(str string) map[string]int {
	counts := make(map[string]int)
	ws, wait := words(str)
	for wrd, ok := <-ws; ok; wrd, ok = <-ws {
		count, exists := counts[wrd]
		if !exists {
			counts[wrd] = 1
		} else {
			counts[wrd] = count + 1
		}

	}
	wait()
	return counts
}

func words(str string) (yield chan string, wait func()) {
	var wg sync.WaitGroup
	wait = wg.Wait
	yield = make(chan string)
	var buffer bytes.Buffer
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, ch := range str {
			if unicode.IsLetter(ch) || unicode.IsDigit(ch) || ch == '_' {
				buffer.WriteRune(ch)
				continue
			}
			if buffer.Len() > 0 {
				yield <- string(buffer.Bytes())
				buffer.Reset()
			}
		}
		if buffer.Len() > 0 {
			yield <- string(buffer.Bytes())
			buffer.Reset()
		}
		close(yield)
	}()
	return yield, wait
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
