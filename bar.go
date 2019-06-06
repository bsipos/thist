package thist

//import "fmt"

func BarSimple(x, y []float64, xlab, ylab []string, title, info string) string {
	if len(xlab) == 0 {
		xlab = AutoLabel(x, Mean(AbsFloats(x)))
	}
	if len(ylab) == 0 {
		ylab = AutoLabel(y, Mean(AbsFloats(y)))
	}
	return Plot(x, y, xlab, ylab, title, info, "#", "@", " ", "_", "|", "-", "|")
}

func Bar(x, y []float64, xlab, ylab []string, title, info string) string {
	if len(xlab) == 0 {
		xlab = AutoLabel(x, Mean(AbsFloats(x)))
	}
	if len(ylab) == 0 {
		ylab = AutoLabel(y, Mean(AbsFloats(y)))
	}
	return Plot(x, y, xlab, ylab, title, info, "\u2588", "\u2591", " ", "_", "\u2502", "\u2500", "\u2524")
}
