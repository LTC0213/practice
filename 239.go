package main

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	var stk []int
	var res []int
	for i := 0; i < len(nums); i++ {
		for len(stk) > 0 && nums[i] > nums[stk[len(stk)-1]] {
			stk = stk[:len(stk)-1]
		}
		stk = append(stk, i)
		for stk[0] < i-k+1 {
			stk = stk[1:]
		}
		if i >= k-1 {
			res = append(res, nums[stk[0]])
		}
	}
	return res
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	res := maxSlidingWindow(nums, 3)
	fmt.Println(res)
	nums = []int{1}
	res = maxSlidingWindow(nums, 1)
	fmt.Println(res)
}
