// run 'go test -coverprofile [reportName]' to generate the coverage report
// run 'go tool cover -html [reportName]' to visualize the report using html

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
		expected := fmt.Sprintf("%.12g", tc.b)
		if res, err := Sqrt(tc.a); err != nil {
			t.Errorf("sqrt(%v) expected %v but got error", tc.a, expected)
		} else {
			if actual := fmt.Sprintf("%.12g", res); actual != expected {
				t.Errorf("sqrt(%v) expected %v but got %v", tc.a, expected, actual)
			}
		}
	}

	failures := []struct {
		a float64
		b string
	}{
		{-2, "cannot Sqrt negative number: -2"},
		{-0.03, "cannot Sqrt negative number: -0.03"},
	}

	for _, tc := range failures {
		if res, err := Sqrt(tc.a); err == nil {
			t.Errorf("sqrt(%v) expected error but got %v", tc.a, fmt.Sprintf("%.12g", res))
		} else {
			if actual := fmt.Sprintf("%v", err); actual != tc.b {
				t.Errorf("sqrt(%v) expected %v but got %v", tc.a, tc.b, actual)
			}
		}
	}
}
