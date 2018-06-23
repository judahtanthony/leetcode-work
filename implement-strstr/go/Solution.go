package main

import "fmt"

func strStr(haystack string, needle string) int {
	findlen, inlen := len(needle), len(haystack)
	switch {
	case findlen == 0:
		return 0
	case inlen > findlen:
		return -1
	case findlen == inlen:
		if haystack == needle {
			return 0
		}
		return -1
	}
	if findlen == 0 {
		return 0
	}
	if inlen == 0 {
		return -1
	}
	to := inlen - findlen
	for i := 0; i <= to; i++ {
		for j := 0; j < findlen; j++ {
			if haystack[i+j] != needle[j] {
				break
			}
			if j == findlen-1 {
				return i
			}
		}
	}

	return -1
}

/**
 * Example 1:
 * Input: haystack = "hello", needle = "ll"
 * Output: 2
 *
 * Example 2:
 * Input: haystack = "aaaaa", needle = "bba"
 * Output: -1
 */
func main() {
	// fmt.Println(strStr("hello", "ll"))
	// fmt.Println(strStr("aaaaa", "bba"))
	fmt.Println(strStr("a", "a"))
}
