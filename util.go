// Copyright © 2019 Botond Sipos
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package thist

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Max(s []float64) float64 {
	if len(s) == 0 {
		return math.NaN()
	}

	max := s[0]
	for _, x := range s {
		if x > max {
			max = x
		}
	}
	return max
}

func Min(s []float64) float64 {
	if len(s) == 0 {
		return math.NaN()
	}

	max := s[0]
	for _, x := range s {
		if x < max {
			max = x
		}
	}
	return max
}

func Mean(s []float64) float64 {
	if len(s) == 0 {
		return math.NaN()
	}

	var sum float64
	for _, x := range s {
		sum += x
	}
	return sum / float64(len(s))
}

func AbsFloats(s []float64) []float64 {
	res := make([]float64, len(s))
	for i, x := range s {
		res[i] = math.Abs(x)
	}
	return res
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func ClearScreen() {
	fmt.Printf("\033[2J")
}

func ClearScreenString() string {
	return "\033[2J"
}

func StringsMaxLen(s []string) int {
	if len(s) == 0 {
		return 0
	}

	max := len(s[0])
	for _, x := range s {
		if len(x) > max {
			max = len(x)
		}
	}
	return max

}

func AutoLabel(s []float64, m float64) []string {
	res := make([]string, len(s))
	nf := false
	var digits int
	if Min(s) < 0 {
		nf = true
	}
	if math.Abs(m) == 0 {
		digits = 5
	} else {
		digits = -int(math.Log10(math.Abs(m) / 5))
	}
	if math.Abs(m) < 10 && digits < 3 {
		digits = 3
	}
	if digits <= 0 {
		digits = 1
	}
	if digits > 8 {
		digits = 8
	}
	dl := 0
	for _, x := range s {
		dg := digits + int(math.Log10(math.Abs(x)+1.0))
		if dg < 0 {
			dg = 0
		}
		if dg > dl {
			dl = dg
		}
	}
	for i, x := range s {
		dg := dl - int(math.Log10(math.Abs(x)+1.0))
		if dg < 0 {
			dg = 0
		}
		f := "%." + strconv.Itoa(dg) + "f"
		ff := f
		if nf && !(x < 0) {
			ff = " " + f
		}
		res[i] = fmt.Sprintf(ff, x)
	}
	return res
}

func RoundFloat64(f float64, n float64) float64 {
	if n == 0.0 {
		return math.Round(f)
	}
	factor := math.Pow(10, float64(n))
	return math.Round(f*factor) / factor
}

// LeftPad2Len https://github.com/DaddyOh/golang-samples/blob/master/pad.go
func LeftPad2Len(s string, padStr string, overallLen int) string {
	var padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = strings.Repeat(padStr, padCountInt) + s
	return retStr[(len(retStr) - overallLen):]
}

// RightPad2Len https://github.com/DaddyOh/golang-samples/blob/master/pad.go
func RightPad2Len(s string, padStr string, overallLen int) string {
	var padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = s + strings.Repeat(padStr, padCountInt)
	return retStr[:overallLen]
}

// https://www.socketloop.com/tutorials/golang-aligning-strings-to-right-left-and-center-with-fill-example
func CenterPad2Len(s string, fill string, n int) string {
	if len(s) >= n {
		return s
	}
	div := (n - len(s)) / 2

	return strings.Repeat(fill, div) + s + strings.Repeat(fill, div)
}
