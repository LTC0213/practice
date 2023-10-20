package main

import "fmt"

func trap(h []int) int {
	if len(h) == 0 {
		return 0
	}
	l, r := 0, len(h)-1
	res := 0
	lmax, rmax := 0, 0
	for l < r {
		lmax = max(lmax, h[l])
		rmax = max(rmax, h[r])
		if h[l] < h[r] {
			res += lmax - h[l]
			l++
		} else {
			res += rmax - h[r]
			r--
		}
	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func trap1(h []int) int {
	n := len(h)
	lmax := make([]int, n)
	lmax[0] = h[0]
	for i := 1; i < n; i++ {
		lmax[i] = max(h[i], lmax[i-1])
	}

	rmax := make([]int, n)
	rmax[n-1] = h[n-1]
	for i := n - 2; i >= 0; i-- {
		rmax[i] = max(h[i], rmax[i+1])
	}

	res := 0
	for i := range h {
		res += min(lmax[i], rmax[i]) - h[i]
	}
	return res
}

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	res := trap(height)
	fmt.Println(res)
}
