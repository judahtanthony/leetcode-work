package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosestBrute(nums []int, target int) int {
	n := len(nums)
	min := math.MaxInt64
	if n < 3 {
		return min
	}
	ret := min
	for x := 0; x < n-2; x++ {
		for y := x + 1; y < n-1; y++ {
			for z := y + 1; z < n; z++ {
				sum := nums[x] + nums[y] + nums[z]
				diff := target - sum
				if diff < 0 {
					diff = -diff
				} else if diff == 0 {
					return sum // Can't get any closer than this.
				}
				if diff < min {
					min = diff
					ret = sum
				}
			}
		}
	}

	return ret
}

func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	min := math.MaxInt64
	if n < 3 {
		return min
	}
	ret := min
	sort.Ints(nums)
	for x := 0; x < n-2; x++ {
		if x > 0 && nums[x] == nums[x-1] {
			continue
		}
		lo, hi := x+1, n-1
		for lo < hi {
			// Only compare unique.
			if lo > x+1 && nums[lo] == nums[lo-1] {
				lo++
				continue
			} else if hi < n-1 && nums[hi] == nums[hi+1] {
				hi--
				continue
			}
			sum := nums[x] + nums[lo] + nums[hi]
			// Can't get any closer than this
			if sum == target {
				return sum
			}
			if sum > target { // Too much
				hi--
			} else { // Not quite there
				lo++
			}
			// Well at least see how close we got.
			diff := target - sum
			if diff < 0 {
				diff = -diff
			}
			if diff < min {
				min = diff
				ret = sum
			}
		}
	}

	return ret
}

func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
	fmt.Println(threeSumClosest([]int{-1, 2, -4}, 1))
	fmt.Println(threeSumClosest([]int{-145, 278, -434545, 1, 65}, 1))
}
