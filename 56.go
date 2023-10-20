package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	var res [][]int

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] || (intervals[i][0] == intervals[j][0] && intervals[i][1] < intervals[j][1])
	})
	prev := intervals[0]
	var cur []int
	for i := 1; i < len(intervals); i++ {
		cur = intervals[i]
		if prev[1] < cur[0] {
			res = append(res, []int{prev[0], prev[1]})
			prev = cur
		} else if cur[1] > prev[1] {
			prev[1] = cur[1]
		}
	}
	res = append(res, prev)
	return res
}

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	res := merge(intervals)
	fmt.Println(res)
}
