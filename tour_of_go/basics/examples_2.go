package basics

import (
	"math"
)

func sqrt(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}

	return lim

}

func pow(x, y float64) float64 {
	if y <= 1 {
		return x
	}

	x = pow(x, y-1) * x
	return x
}

func pow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}
