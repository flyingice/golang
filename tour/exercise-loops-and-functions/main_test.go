package exercise

import (
	"fmt"
	"math"
	"testing"
)

func TestSqrt(t *testing.T) {
	tests := []struct{ a, b float64 }{
		// edge cases
		// fail {.0, math.Sqrt(.0)},
		{math.SmallestNonzeroFloat64, math.Sqrt(math.SmallestNonzeroFloat64)},
		{math.SmallestNonzeroFloat32, math.Sqrt(math.SmallestNonzeroFloat32)},
		{.0001, math.Sqrt(.0001)},
		{.1, math.Sqrt(.1)},
		{1, math.Sqrt(1)},
		{math.MaxFloat32, math.Sqrt(math.MaxFloat32)},
		{math.MaxFloat64, math.Sqrt(math.MaxFloat64)},
		// normal cases
		{59.8, math.Sqrt(59.8)},
		{225, math.Sqrt(225)},
		{250.369, math.Sqrt(250.369)},
		{1024, math.Sqrt(1024)},
	}

	for _, tc := range tests {
		if actual, expected := fmt.Sprintf("%.12g", Sqrt(tc.a)), fmt.Sprintf("%.12g", tc.b); actual != expected {
			t.Errorf("sqrt(%v) expected %v but got %v", tc.a, expected, actual)
		}
	}
}
