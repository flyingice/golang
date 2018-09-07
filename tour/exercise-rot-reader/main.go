package exercise

import (
	"io"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
	if b >= 'A' && b <= 'M' || b >= 'a' && b <= 'm' {
		return b + 13
	} else if b >= 'N' && b <= 'Z' || b >= 'n' && b <= 'z' {
		return b - 13
	} else {
		return b
	}
}

func (rr *rot13Reader) Read(buffer []byte) (int, error) {
	n, err := rr.r.Read(buffer)
	for i, v := range buffer {
		buffer[i] = rot13(v)
	}
	return n, err
}
