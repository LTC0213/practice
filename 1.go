package main

import "fmt"

func twoSum(nums []int, target int) []int {
	book := make(map[int]int)
	var res []int
	for i, num := range nums {
		j, ok := book[target-num]
		if ok {
			res = append(res, i, j)
		}
		book[num] = i
	}
	return res
}

func main() {
	nums := []int{2, 7, 11, 15}
	res := twoSum(nums, 9)
	fmt.Println(res)
}
