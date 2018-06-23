package main

import "fmt"

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	c, i := 0, 1
	for i < n {
		if nums[c] != nums[i] {
			c++
			nums[c] = nums[i]
		}
		i++
	}

	return c + 1
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	// nums is passed in by reference. (i.e., without making a copy)
	len := removeDuplicates(nums)

	// any modification to nums in your function would be known by the caller.
	// using the length returned by your function, it prints the first len elements.
	for i := 0; i < len; i++ {
		fmt.Println(nums[i])
	}
}
