package main

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/tour/pic"
)

// IPAddr represents a 4-byte IP address
type IPAddr [4]byte

// ErrNegativeSqrt is raised by trying to take the square root of a negative number.
type ErrNegativeSqrt float64

func main() {
	exercise01()
	exercise02()
	exercise03()
	exercise04()
	exercise05()
	exercise06()
	exercise07()
}

// MyReader - Reads an infinite number of 'A' bytes
type MyReader struct{}

func (r MyReader) Read(p []byte) (int, error) {
	for i := range p {
		p[i] = 'A'
	}
	return len(p), nil
}

func exercise07() {

}

func exercise06() {
	Sqrt(-2, retry10Times)
}

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("cannot Sqrt negative number: ", float64(e))
}

func exercise05() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

func (ip IPAddr) String() string {
	strs := []string{
		byteToString(ip[0]),
		byteToString(ip[1]),
		byteToString(ip[2]),
		byteToString(ip[3])}
	return strings.Join(strs, ".")
}

func iter(anys []interface{}) func() (interface{}, bool) {
	i := 0
	return func() (interface{}, bool) {
		for ; i < len(anys); i++ {
			return anys[i], true
		}
		return nil, false
	}
}

func byteToString(b byte) string {
	return strconv.Itoa(int(b))
}

func exercise04() {
	f := fibonacci()
	for i := int64(0); i < 50; i++ {
		fmt.Println(f(i))
	}
}

func fibonacci() func(int64) int64 {
	cache := make(map[int64]int64)
	cache[0] = 0
	cache[1] = 1
	return func(n int64) int64 {
		result := int64(0)
		defer func() { cache[n] = result; n++ }()
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
	iterator := wordIterator(str)
	for wrd, ok := iterator(); ok; wrd, ok = iterator() {
		count, exists := counts[wrd]
		if !exists {
			counts[wrd] = 1
		} else {
			counts[wrd] = count + 1
		}
	}
	return counts
}

func wordIterator(str string) func() (string, bool) {
	var buffer bytes.Buffer
	i := 0
	var runes = []rune(str)
	return func() (string, bool) {
		for ; i < len(runes); i++ {
			ch := runes[i]
			if unicode.IsLetter(ch) || unicode.IsDigit(ch) || ch == '_' {
				buffer.WriteRune(ch)
			} else if buffer.Len() > 0 {
				token := string(buffer.Bytes())
				buffer.Reset()
				return token, true
			}
		}
		if buffer.Len() > 0 {
			token := string(buffer.Bytes())
			buffer.Reset()
			return token, true
		}
		return "", false
	}
}

func exercise02() {
	pic.Show(GraphExtension(10, 10, func(x, y int) uint8 {
		return (uint8(x) + uint8(y)) / uint8(2)
	}))
}

// GraphExtension graphs an arbitrary x,y function
func GraphExtension(width, height int, fn func(x, y int) uint8) func(int, int) [][]uint8 {
	return func(width, height int) [][]uint8 {
		ys := make([][]uint8, height)
		for y := 0; y < height; y++ {
			xs := make([]uint8, width)
			for x := 0; x < width; x++ {
				xs = append(xs, fn(x, y))
			}
			ys[y] = xs
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
func Sqrt(x float64, terminate func(int, float64, float64) bool) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	var z0, z1 float64
	z0 = x
	z1 = z0
	for i := 0; i < 1 || !terminate(i, z0, z1); i++ {
		z0 = z1
		z1 = newton(z0, x)
	}
	return z1, nil
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
