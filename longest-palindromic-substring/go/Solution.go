package main

import "fmt"

func checkPal(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func longestPalindrome(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	max, lo, hi := string(s[0]), 0, 1
	for hi <= n {
		check := s[lo:hi]
		if checkPal(check) {
			max = check
		}
		hi++
		if hi > n {
			lo++
			hi = lo + len(max) + 1
		}
	}
	return max
}

func main() {
	fmt.Println(longestPalindrome("abfdgabcdefg0123456789876543210gfedcbadrjfsahcbhfbklhmnfaguigjrskleafbjdahgdrklradghdfjdd"))
}
