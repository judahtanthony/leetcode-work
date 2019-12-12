package main

import "fmt"

func characterReplacement(s string, k int) int {
	n := len(s)
	if n <= k {
		return n
	}
	maxLength := k
	var remaining int
	for c := 0; c < n; c++ {
		if c > 0 && s[c] == s[c-1] {
			continue
		}
		remaining = k
		l, r := c, c
		for r < n {
			if s[r] == s[c] {
				r++
			} else if remaining > 0 {
				remaining--
				r++
			} else {
				break
			}
		}
		for l > 0 {
			if s[l-1] == s[c] {
				l--
			} else if remaining > 0 {
				remaining--
				l--
			} else {
				break
			}
		}
		if r-l > maxLength {
			maxLength = r - l
		}
	}

	return maxLength
}

func main() {
	fmt.Println(characterReplacement("AABABBA", 2))
}
