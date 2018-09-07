package exercise

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func equal(x, y, epsilon float64) bool {
	return math.Abs(x-y)/math.Max(x, y) < epsilon
}

func Sqrt(x float64) (float64, error) {
	if x < .0 {
		return .0, ErrNegativeSqrt(x)
	}

	guess := 1.
	for {
		// overflow with next := guess - (guess*guess-x)/(2*guess)
		next := guess + (x/guess-guess)/2.
		if equal(next, guess, 1e-12) {
			break
		}
		guess = next
	}

	return guess, nil
}
