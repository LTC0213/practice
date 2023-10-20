package main

import "fmt"

func productExceptSelf(nums []int) []int {
	n := len(nums)
	L := make([]int, n)
	R := make([]int, n)
	L[0] = nums[0]
	for i := 1; i < n; i++ {
		L[i] = L[i-1] * nums[i]
	}
	R[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		R[i] = R[i+1] * nums[i]
	}

	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = 1
		if i > 0 {
			res[i] *= L[i-1]
		}
		if i < n-1 {
			res[i] *= R[i+1]
		}
	}
	return res
}

func main() {
	nums := []int{1, 2, 3, 4}
	res := productExceptSelf(nums)
	fmt.Println(res)
}
