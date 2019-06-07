package thist

import (
	"fmt"
	terminal "github.com/wayneashleyberry/terminal-dimensions"
	"strconv"
	"strings"
)

func Plot(x, y []float64, xlab, ylab []string, title string, info []string, symbol, negSymbol, space, top, vbar, hbar, tvbar string) string {
	if len(x) == 0 {
		return ""
	}
	// Based on: http://pyinsci.blogspot.com/2009/10/ascii-histograms.html
	width, _ := terminal.Width()
	height, _ := terminal.Height()

	xll := StringsMaxLen(xlab)
	yll := StringsMaxLen(ylab)
	width -= uint(yll + 1)

	res := strings.Repeat(space, yll+1) + CenterPad2Len(title, space, int(width)) + "\n"
	height -= 4
	height -= uint(len(info))

	height -= uint(xll + 1)

	xf := xFactor(len(x), int(width))
	if xf < 1 {
		xf = 1
	}

	if xll < xf-2 {
		height += uint(xll - 1)
	}

	ny := normalizeY(y, int(height))

	block := strings.Repeat(symbol, xf)
	nblock := strings.Repeat(negSymbol, xf)
	if xf > 2 {
		block = vbar + strings.Repeat(symbol, xf-2) + vbar
		nblock = vbar + strings.Repeat(negSymbol, xf-2) + vbar
	}

	blank := strings.Repeat(space, xf)
	topBar := strings.Repeat(top, xf)

	for l := int(height); l > 0; l-- {
		if yll > 0 {
			found := false
			for i, t := range ny {
				if l == t {
					res += fmt.Sprintf("%-"+strconv.Itoa(yll)+"s"+tvbar, ylab[i])
					found = true
					break
				}
			}
			if !found {
				res += strings.Repeat(space, yll) + vbar
			}
		}
		for _, c := range ny {
			if l == Abs(c) {
				res += topBar
			} else if l < Abs(c) {
				if c < 0 {
					res += nblock
				} else {
					res += block
				}
			} else {
				res += blank
			}
		}
		res += "\n"
	}

	if xll > 0 {
		res += strings.Repeat(space, yll) + vbar + strings.Repeat(hbar, int(width)) + "\n"
		if xll < xf-2 {
			res += strings.Repeat(space, yll) + vbar
			for _, xl := range xlab {
				res += vbar + RightPad2Len(xl, space, xf-1)
			}
		} else {
			for i := 0; i < xll; i++ {
				res += strings.Repeat(space, yll) + vbar
				for j := yll + 1; j < int(width); j++ {
					if (j-yll-1)%xf == 0 {
						bin := (j - yll - 1) / xf
						if bin < len(xlab) && i < len(xlab[bin]) {
							res += string(xlab[bin][i])
						} else {
							res += space
						}
					} else {
						res += space
					}
				}
				res += "\n"
			}

		}
	}

	for _, il := range info {
		res += il + "\n"
	}
	return res
}

func normalizeY(y []float64, height int) []int {
	max := Max(y)
	res := make([]int, len(y))

	for i, x := range y {
		res[i] = int(x / max * float64(height))
	}
	return res
}

func xFactor(n int, width int) int {
	return int(width / n)
}
