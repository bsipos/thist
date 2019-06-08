// Copyright © 2016 Wei Shen <shenwei356@gmail.com>
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
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func (h *Hist) SaveImage(f string) {
	data := plotter.Values(h.Counts)

	if h.Normalize {
		data = plotter.Values(h.NormCounts())
	}

	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = h.Title
	p.Y.Label.Text = "Count"
	if h.Normalize {
		p.Y.Label.Text = "Frequency"
	}

	w := vg.Points(20)

	bars, err := plotter.NewBarChart(data, w)
	if err != nil {
		panic(err)
	}
	bars.LineStyle.Width = vg.Length(0)
	bars.Color = plotutil.Color(2)
	bars.Offset = w

	p.Add(bars)

	xlab := make([]string, len(h.BinStart))
	for i, bin := range h.BinStart {
		xlab[i] = fmt.Sprintf("%.3f", bin)
	}

	p.NominalX(xlab...)

	if err := p.Save(5*vg.Inch, 3*vg.Inch, f); err != nil {
		panic(err)
	}
}
