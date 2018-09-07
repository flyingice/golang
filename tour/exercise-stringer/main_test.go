package exercise

import (
	"fmt"
	"testing"
)

func TestIPAddrStringer(t *testing.T) {
	tests := []struct {
		a IPAddr
		b string
	}{
		{IPAddr{127, 0, 0, 1}, "127.0.0.1"},
		{IPAddr{8, 8, 8, 8}, "8.8.8.8"},
		{IPAddr{192, 168, 0, 1}, "192.168.0.1"},
	}

	for _, tc := range tests {
		if actual, expected := fmt.Sprintf("%v", tc.a), tc.b; actual != expected {
			t.Errorf("expected %v while got %v\n", expected, actual)
		}
	}
}
