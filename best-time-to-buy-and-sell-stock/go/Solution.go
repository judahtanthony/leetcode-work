package main

import "fmt"

func maxProfit(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	minValue := prices[0]
	maxValue := prices[0]
	maxProfit := 0

	for _, v := range prices {
		if v < minValue {
			minValue = v
			maxValue = v
		}
		if v > maxValue {
			maxValue = v
		} else {
			continue
		}
		profit := maxValue - minValue
		if profit > maxProfit {
			maxProfit = profit
		}
	}

	return maxProfit
}

func main() {
	fmt.Print(maxProfit([]int{5, 6, 7, 10, 0, 2, 6}))
}
