/*
Exercise 7.5: The LimitReader function in the io package accepts an io.Reader r
and a number of bytes n, and returns another Reader that reads from r but reports
an end-of-file conditionafternbytes. Implement it.
func LimitReader(r io.Reader, n int64) io.Reader
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type ReaderWithLimit struct {
	reader io.Reader
	limit  int64
}

func (r *ReaderWithLimit) Read(buffer []byte) (int, error) {
  if r.limit <= 0 {
		return 0, io.EOF
	}

  if r.limit < int64(len(buffer)) {
		buffer = buffer[:r.limit]
	}

  n, err := r.reader.Read(buffer)
	r.limit -= int64(n)

  return n, err
}

func LimitReader(r io.Reader, n int64) io.Reader {
	lr := &ReaderWithLimit{r, n}
	return lr
}

func main() {
	const input = "Hello, rainforest\nHello, again"
	scanner := bufio.NewScanner(LimitReader(strings.NewReader(input), 11))
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

