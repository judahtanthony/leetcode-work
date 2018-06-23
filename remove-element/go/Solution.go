package main

import "fmt"

func removeElement(nums []int, val int) int {
	n := len(nums)
	if n == 0 {
		return n
	}
	slow, fast := 0, 0
	for fast < n {
		if nums[fast] == val {
			fast++
			continue
		}
		if fast != slow {
			nums[slow] = nums[fast]
		}
		fast++
		slow++
	}

	return slow
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	// nums is passed in by reference. (i.e., without making a copy)
	len := removeElement(nums, 4)

	// any modification to nums in your function would be known by the caller.
	// using the length returned by your function, it prints the first len elements.
	for i := 0; i < len; i++ {
		fmt.Println(nums[i])
	}
}
