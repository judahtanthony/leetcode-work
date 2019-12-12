package main

import "fmt"

func climbStairs(n int) int {
	v := climbStairsDynamic(n)
	//v := climbStairsBottomUp(n)
	return v
}

func climbStairsBottomUp(n int) int {
	if n <= 3 {
		return n
	}

	cache := make([]int, n+1, n+1)
	cache[0] = 0
	cache[1] = 1
	cache[2] = 2
	cache[3] = 3

	for i := 4; i <= n; i++ {
		cache[i] = cache[i-1] + cache[i-2]
	}

	return cache[n]
}

func climbStairsDynamic(n int) int {
	cache := make([]int, n+1, n+1)
	for i := 0; i <= n; i++ {
		if i <= 3 {
			cache[i] = i
		} else {
			cache[i] = 0
		}
	}
	return climbStairsWithCache(n, cache)
}

func climbStairsWithCache(n int, cache []int) int {
	if n <= 2 {
		return n
	}

	if cache[n] > 0 {
		return cache[n]
	}

	v := climbStairs(n-1) + climbStairs(n-2)
	cache[n] = v

	return v
}

func main() {
	fmt.Println(climbStairs(6))
}
