package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	var res [][]int
	n := len(nums)
	sort.Ints(nums)
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j, k := i+1, n-1; j < k; {
			if j > i+1 && nums[j] == nums[j-1] {
				j++
				continue
			}
			sum := nums[i] + nums[j] + nums[k]
			if sum <= 0 {
				if sum == 0 {
					res = append(res, []int{nums[i], nums[j], nums[k]})
				}
				j++
				continue
			} else {
				k--
			}
		}
	}
	return res
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	res := threeSum(nums)
	fmt.Println(res)
}
