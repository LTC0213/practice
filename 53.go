package main

import "math"

func maxSubArray(nums []int) int {
	sum, res := 0, math.MinInt32
	for _, num := range nums {
		sum += num
		res = max(res, sum)
		sum = max(sum, 0)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
