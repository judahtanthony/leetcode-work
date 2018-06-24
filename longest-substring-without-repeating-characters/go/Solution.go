package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n < 2 {
		return n
	}
	max := 1
	curr := 0
	m := make(map[byte]int)
	for i := 0; i < n; i++ {
		// Have we seen this already?
		if oldi, ok := m[s[i]]; ok {
			// remove old chars
			for i-curr <= oldi {
				delete(m, s[i-curr])
				curr--
			}
		}
		// Unique
		m[s[i]] = i
		curr++
		if curr > max {
			max = curr
		}
	}
	return max
}

func main() {
	fmt.Println(lengthOfLongestSubstring("dvdf"))
	fmt.Println(lengthOfLongestSubstring("aab"))
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
