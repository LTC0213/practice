package main

import "fmt"

func maxArea(h []int) int {
	if len(h) == 0 {
		return 0
	}
	l, r := 0, len(h)-1
	res := 0
	for l < r {
		area := min(h[l], h[r]) * (r - l)
		res = max(res, area)
		if h[l] < h[r] {
			l++
		} else {
			r--
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	h := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	res := maxArea(h)
	fmt.Println(res)
}
