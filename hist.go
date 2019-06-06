package thist

import (
	"fmt"
	"math"
	"sort"
)

type Hist struct {
	Title     string
	Info      string
	BinMode   string
	MaxBins   int
	NrBins    int
	DataCount int
	DataMin   float64
	DataMax   float64
	DataMean  float64
	DataSd    float64
	Normalize bool
	BinStart  []float64
	BinEnd    []float64
	Counts    []float64
	m         float64
	Precision float64
}

func NewHist(data []float64, title, info, binMode string, maxBins int, normalize bool) *Hist {
	h := &Hist{title, info, binMode, maxBins, 0, 0, math.NaN(), math.NaN(), math.NaN(), math.NaN(), normalize, []float64{}, []float64{}, []float64{}, math.NaN(), math.Pow(1, -8)}
	if h.BinMode == "" {
		h.BinMode = "fit"
	}

	for _, d := range data {
		h.Update(d)
	}
	fmt.Println(h.Dump())

	return h
}

func (h *Hist) buildBins() ([]float64, []float64, float64) {
	var n int
	var w float64

	if h.DataMin == h.DataMax {
		n = 1
		w = 1
	} else if h.BinMode == "fixed" {
		n = h.MaxBins
		w = (h.DataMax - h.DataMin) / float64(n)
	} else if h.BinMode == "auto" || h.BinMode == "fit" {
		w = scottsRule(h.DataCount, h.DataSd)
		n = int((h.DataMax - h.DataMin) / w)
		if n < 1 {
			n = 1
		}
	}

	s := make([]float64, n)
	e := make([]float64, n)

	for i := 0; i < n; i++ {
		s[i] = h.DataMin + float64(i)*w
		e[i] = h.DataMin + float64(i+1)*w
	}

	return s, e, w

}

func (h *Hist) updateMoments(p float64) {
	oldMean := h.DataMean
	h.DataMean += (p - h.DataMean) / float64(h.DataCount)
	h.m += (p - oldMean) * (p - h.DataMean)
	h.DataSd = math.Sqrt(h.m / float64(h.DataCount))
}

func scottsRule(n int, sd float64) float64 {
	h := (3.5 * sd) / math.Pow(float64(n), 1.0/3.0)
	return h
}

func (h *Hist) Update(p float64) {
	h.DataCount++
	oldMin := h.DataMin
	oldMax := h.DataMax
	if math.IsNaN(h.DataMin) || p < h.DataMin {
		h.DataMin = p
	}
	if math.IsNaN(h.DataMax) || p > h.DataMax {
		h.DataMax = p
	}
	if h.DataCount == 1 {
		h.DataMean = p
		h.DataSd = 0.0
		h.m = 0.0
		h.buildBins()
		bs, be, _ := h.buildBins()
		h.Counts = []float64{1.0}
		h.BinStart = bs
		h.BinEnd = be
		return
	} else {
		h.updateMoments(p)
	}

	if !math.IsNaN(oldMin) && p >= oldMin && !math.IsNaN(oldMax) && p <= oldMax {
		var i int
		if p == oldMin {
			i = 0
		} else if p == oldMax {
			i = len(h.Counts) - 1
		} else {
			i = sort.SearchFloat64s(h.BinStart, p) - 1
			if i < 0 {
				i = 0
			}
		}
		h.Counts[i]++
		return
	}

	bs, be, bw := h.buildBins()
	_ = bw
	newCounts := make([]float64, len(bs))

	h.BinStart = bs
	h.BinEnd = be
	h.Counts = newCounts

}

func (h *Hist) Draw() string {
	return ""
}

func (h *Hist) Summary() string {
	res := ""
	return res
}

func (h *Hist) Dump() string {
	res := "Bin\tBinStart\tBinEnd\tCount\n"

	for i, c := range h.Counts {
		res += fmt.Sprintf("%d\t%.4f\t%.4f\t%.0f\n", i, h.BinStart[i], h.BinEnd[i], c)
	}

	return res
}
