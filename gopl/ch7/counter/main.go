// Exercise 7.1: Using the ideas from ByteCounter, implement counters for words and for lines.

package main

import (
	"bufio"
	"fmt"
	"bytes"
)

type LineCounter int
type WordCounter int

func (lc *LineCounter) Write(text []byte) (int, error) {
	s := bufio.NewScanner(bytes.NewReader(text))
	for s.Scan() {
		*lc += LineCounter(1)
	}
	return len(text), nil
}

func (wc *WordCounter) Write(text []byte) (int, error) {
	s := bufio.NewScanner(bytes.NewReader(text))
	s.Split(bufio.ScanWords)
	for s.Scan() {
		*wc += WordCounter(1)
	}
	return len(text), nil
}

func main() {
	const input = "hello\nrainforest\nfrom test"

	var lc LineCounter
	fmt.Fprintf(&lc, "%s", input)
	fmt.Println(lc)

	var wc WordCounter
	fmt.Fprintf(&wc, "%s", input)
	fmt.Println(wc)
}
