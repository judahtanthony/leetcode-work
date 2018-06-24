package main

import "fmt"

func convert(s string, numRows int) string {
	n := len(s)
	if numRows == 1 || n == 1 {
		return s
	}
	max := numRows*2 - 2
	ret := make([]byte, n)
	retlen := 0
	for i := 0; i < numRows; i++ {
		also := max - (2 * i)
		for at := i; at < n; at += max {
			ret[retlen] = s[at]
			retlen++
			if also > 0 && also < max && at+also < n {
				ret[retlen] = s[at+also]
				retlen++
			}
		}
	}
	return string(ret)
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
	fmt.Println(convert("PAYPALISHIRING", 4))
}
