package main

import "fmt"

func twoSum(nums []int, target int) []int {
	ret := make([]int, 2)
	cache := make(map[int]int)
	for k, v := range nums {
		if f, ok := cache[v]; ok {
			ret[0] = f
			ret[1] = k
			break
		}
		cache[target-v] = k
	}
	return ret
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
