package main

import "fmt"

func getSum(a int, b int) int {
	result := 0
	overflow := 0

	for {
		overflow = a & b
		result = a ^ b
		a = overflow << 1
		b = result

		if overflow == 0 {
			break
		}
	}

	return result
}

func main() {
	fmt.Print(getSum(-2, 3))
}
