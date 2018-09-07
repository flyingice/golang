package exercise

import (
	"bufio"
	"strings"
	"testing"
)

func TestRotReader(t *testing.T) {
	tests := []struct{ a, b string }{
		{"Lbh penpxrq gur pbqr!", "You cracked the code!"},
	}

	for _, tc := range tests {
		scanner := bufio.NewScanner(&rot13Reader{strings.NewReader(tc.a)})
		expected := tc.b
		if scanner.Scan() {
			if actual := scanner.Text(); actual != expected {
				t.Errorf("encrypt: %v expected: %v while got: %v", tc.a, expected, actual)
			}
		} else {
			t.Errorf("encrypt: %v expected: %v while got an error", tc.a, expected)
		}
	}
}
