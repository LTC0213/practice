package main

import "fmt"

func moveZeroes(nums []int) {
	for i, j := 0, 0; j < len(nums); j++ {
		if nums[j] == 0 {
			continue
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}
