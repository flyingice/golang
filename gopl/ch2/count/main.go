package main

import (
	"fmt"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i / 2] + byte(i & 1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x >> (8 * 0))] +
		pc[byte(x >> (8 * 1))] +
		pc[byte(x >> (8 * 2))] +
		pc[byte(x >> (8 * 3))] +
		pc[byte(x >> (8 * 4))] +
		pc[byte(x >> (8 * 5))] +
		pc[byte(x >> (8 * 6))] +
		pc[byte(x >> (8 * 7))])
}

/*
Exercise 2.3: Rewrite PopCount to use a loop instead of a single expression.
*/
func PopCountIterative(x uint64) int {
	var cnt int
	for i := 0; i < 8; i++ {
		cnt += int(pc[byte(x)])
		x >>= 8
	}
	
	return cnt
}

/*
Exercise 2.4: Write a version of PopCount that counts bits by shifting its argument through 64 bit positions, 
testing the rightmost bit each time.
*/
func PopCountTinyStep(x uint64) int {
	var cnt int;
	for i := 0; i < 64; i++ {
		if x & 1 == 1 {
			cnt++
		}
		x >>= 1
	}
	
	return cnt
}

/*
Exercise 2.5: The expression x&(x-1) clears the rightmost non-zero bit of x. 
Write a version of PopCount that counts bits by using this fact.
*/
func PopCountSmartClear(x uint64) int {
	var cnt int;
	for x > 0 {
		x = x & (x - 1)
		cnt++
	}
	
	return cnt
}

func main() {
	var i uint64 = 255
	fmt.Printf("%v has %v bit(s) set.\n", i, PopCountSmartClear(i))
}
