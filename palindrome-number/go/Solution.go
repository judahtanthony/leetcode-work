package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	end := 10
	for x/end >= 10 {
		end = end * 10
	}
	for {
		if x%10 != (x/end)%10 {
			return false
		}
		if x < 10 || end < 100 {
			break
		}
		x /= 10
		end /= 100
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(121))
}
