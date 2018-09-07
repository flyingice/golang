package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	counter := make(map[string]int)
	for _, word := range words {
		if _, ok := counter[word]; ok {
			counter[word]++
		} else {
			counter[word] = 1
		}
	}

	return counter
}

func main() {
	wc.Test(WordCount)
}
