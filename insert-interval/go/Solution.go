package main

import "fmt"

func insert(intervals [][]int, newInterval []int) [][]int {
	result := make([][]int, 0)

	result = append(result, []int{newInterval[0], newInterval[1]})
	ri := 0
	for _, r := range intervals {
		if r[1] < result[ri][0] { // prepend
			result = append(result, []int{r[0], r[1]})
			ri++
			result[ri-1], result[ri] = result[ri], result[ri-1]
		} else if r[0] > result[ri][1] { // append
			result = append(result, []int{r[0], r[1]})
			ri++
		} else { // merge
			if r[0] < result[ri][0] {
				result[ri][0] = r[0]
			}
			if r[1] > result[ri][1] {
				result[ri][1] = r[1]
			}
		}
	}

	return result
}

func main() {
	fmt.Println(insert([][]int{{1, 3}, {6, 9}}, []int{2, 5}))
}
