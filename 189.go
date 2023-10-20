package main

import "fmt"

func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	var reverse func(i, j int)
	reverse = func(i, j int) {
		for i < j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}
	reverse(0, n-1)
	reverse(0, k-1)
	reverse(k, n-1)
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(nums, k)
	fmt.Println(nums)
}
