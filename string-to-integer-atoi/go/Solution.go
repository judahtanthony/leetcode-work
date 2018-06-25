package main

import "fmt"

const INT_MAX = (1 << 31) - 1
const INT_MIN = 1 << 31

func myAtoi(str string) int {
	n := len(str)
	if n == 0 {
		return 0
	}
	val, curr, neg := 0, 0, false
	// get to number
	for str[curr] == ' ' {
		curr++
		if curr >= n {
			return 0
		}
	}
	// consume sign
	if str[curr] == '-' {
		neg = true
		curr++
	} else if str[curr] == '+' {
		curr++
	}
	for curr < n && str[curr] >= '0' && str[curr] <= '9' {
		v := int(str[curr] - '0')
		if !neg && (INT_MAX-v)/10 < val {
			return INT_MAX
		} else if neg && (INT_MIN-v)/10 < val {
			return -INT_MIN // Go won't let me overflow
		}
		val = val*10 + v
		curr++
	}
	if neg {
		val *= -1
	}
	return val
}

func main() {
	fmt.Println(myAtoi("42"))
	fmt.Println(myAtoi("-42"))
	fmt.Println(myAtoi("4193 with words"))
	fmt.Println(myAtoi("words and 987"))
	fmt.Println(myAtoi("-91283472332"))
	fmt.Println(myAtoi(" "))
	fmt.Println(myAtoi("2147483648"))
	fmt.Println(myAtoi("-2147483649"))
	fmt.Println()
	fmt.Println(INT_MAX)
	fmt.Println(INT_MIN)
}
