/*
Exercise 7.2: Write a function CountingWriter with the signature below that, 
given an io.Writer, returns a new Writer that wraps the original, and a pointer 
to an int64 variable that at any moment contains the number of bytes written to 
the new Writer.
func CountingWriter(w io.Writer) (io.Writer, *int64)
*/

package main

import (
	"fmt"
	"io"
	"os"
)

type CounterWriter struct {
	w io.Writer
	c int64
}

func (cw *CounterWriter) Write(buffer []byte) (int, error) {
	n, err := cw.w.Write(buffer)
	cw.c += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := &CounterWriter{w, 0}
	return cw, &cw.c
}

func main() {
	w, n := CountingWriter(os.Stdout)

	fmt.Fprintf(w, "hello, word!\n")
	fmt.Printf("writer count [%d]\n", *n)
	fmt.Fprintf(w, "1234567890\n")
	fmt.Printf("writer count [%d]\n", *n)

	fmt.Printf("%T %v", w, w)
}
