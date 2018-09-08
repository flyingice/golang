package main

func lengthOfLongestSubstring(s string) int {
	var maxLength, start int
	m := make(map[byte]int)
	for i, c := range []byte(s) {
		if v, ok := m[c]; ok {
			for j := start; j < v; j++ {
				delete(m, s[j])
			}
			start = v + 1
		} else {
			if i-start+1 > maxLength {
				maxLength = i - start + 1
			}
		}

		m[c] = i
	}

	return maxLength
}

func main() {
}
