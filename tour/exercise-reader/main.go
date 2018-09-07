package main

import (
	"golang.org/x/tour/reader"
)

type MyReader struct{}

func (MyReader) Read(buffer []byte) (int, error) {
	buffer = buffer[:cap(buffer)]
	for i := range buffer {
		buffer[i] = 'A'
	}
	return cap(buffer), nil
}

func main() {
	reader.Validate(MyReader{})
}
