package exercise

import (
	"math"
)

func equal(x, y, epsilon float64) bool {
	return math.Abs(x-y)/math.Max(x, y) < epsilon
}

func Sqrt(x float64) float64 {
	guess := 1.
	for {
		// overflow with next := guess - (guess*guess-x)/(2*guess)
		next := guess + (x/guess-guess)/2.
		if equal(next, guess, 1e-12) {
			break
		}
		guess = next
	}

	return guess
}
