package main

import (
  "fmt"
  "os"
  "strings"
  "time"
)

func main() {
  // Exercise 1.1: Modify the echo program to also print os.Args[0]
  start := time.Now()
  fmt.Println(strings.Join(os.Args, " "))
  fmt.Printf("Elapsed time: %g\n", time.Since(start).Seconds())

  // Exercise 1.2: Modify the echo program to print the index and value
  // of each of its arguments, one per line.
  for i, v := range os.Args {
    fmt.Println(i, v)
  }

  // Exercise 1.3: Experiment to measure the difference in running time between
  // our potentially inefficient versions and the one that uses strings.Join.
  start = time.Now()
  s, sep := "", ""
  for _, v := range os.Args {
    s += v + sep
    sep = " "
  }
  fmt.Println(s)
  fmt.Printf("Elapsed time: %g\n", time.Since(start).Seconds())
}
