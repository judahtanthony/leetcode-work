package main

import (
	"fmt"
	"sort"
)

func threeSumBrute(nums []int) [][]int {
	r := make([][]int, 0)
	n := len(nums)
	if n > 2 {
		sort.Ints(nums)
		for a := 0; a < n-2; a++ {
			if a > 0 && nums[a] == nums[a-1] {
				continue
			}
			sum := nums[a]
			if sum > 0 {
				break
			}
			for b := a + 1; b < n-1; b++ {
				if b > a+1 && nums[b] == nums[b-1] {
					continue
				}
				sum = nums[a] + nums[b]
				if sum > 0 {
					break
				}
				for c := b + 1; c < n; c++ {
					if c > b+1 && nums[c] == nums[c-1] {
						continue
					}
					sum = nums[a] + nums[b] + nums[c]
					if sum > 0 {
						break
					}
					if sum == 0 {
						r = append(r, []int{nums[a], nums[b], nums[c]})
					}
				}
			}
		}
	}
	return r
}

func threeSum(nums []int) [][]int {
	r := make([][]int, 0)
	n := len(nums)
	if n > 2 {
		sort.Ints(nums) // Can this be done more efficiently without sorting and the > 0 cutoff?
		// Dynamic programming for the win
		lookup := make(map[int]int, n)
		for v, k := range nums {
			lookup[k] = v
		}
		for a := 0; a < n-2; a++ {
			if a > 0 && nums[a] == nums[a-1] {
				continue
			}
			sum := nums[a]
			if sum > 0 {
				break
			}
			for b := a + 1; b < n-1; b++ {
				if b > a+1 && nums[b] == nums[b-1] {
					continue
				}
				sum = nums[a] + nums[b]
				if sum > 0 {
					break
				}
				if v, ok := lookup[-sum]; ok && v > b {
					r = append(r, []int{nums[a], nums[b], -sum})
				}
			}
		}
	}
	return r
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1}))
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{0, 0, 0, 0}))
	fmt.Println(threeSum([]int{1, 2, -2, -1}))
	fmt.Println(threeSum([]int{-2, 0, 1, 1, 2}))
}
