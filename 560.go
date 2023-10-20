package main

import "fmt"

func subarraySum(nums []int, k int) int {
	sum, res := 0, 0
	hash := make(map[int]int)
	hash[0] = 1
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		res += hash[sum-k]
		hash[sum]++
	}
	return res
}

func main() {
	nums := []int{1, 1, 1}
	res := subarraySum(nums, 2)
	fmt.Println(res)
	nums = []int{1, 2, 3}
	res = subarraySum(nums, 3)
	fmt.Println(res)
}
