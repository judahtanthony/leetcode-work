package main

import "fmt"

// OPTIMIZE!  This can be done in linear time.
func maxAreaBruteForce(height []int) int {
	n := len(height)
	// assert(n > 1)
	max := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			y := height[i]
			if height[j] < y {
				y = height[j]
			}
			a := y * (j - i)
			if a > max {
				max = a
			}
		}
	}
	return max
}

func maxArea(height []int) int {
	// assert(n > 1)
	max, left, right := 0, 0, len(height)-1
	dist := right
	for left < right {
		if height[left] < height[right] {
			if a := height[left] * dist; a > max {
				max = a
			}
			left++
		} else {
			if a := height[right] * dist; a > max {
				max = a
			}
			right--
		}
		dist--
	}
	return max
}

func main() {
	fmt.Println(maxArea([]int{1, 1}))
	fmt.Println(maxArea([]int{1, 2}))
	fmt.Println(maxArea([]int{10, 10, 10}))
}
