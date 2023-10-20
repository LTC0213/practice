package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	for i, j := 0, n-1; i < m && j >= 0; {
		cur := matrix[i][j]
		if cur == target {
			return true
		}
		if cur < target {
			i++
		} else {
			j--
		}
	}
	return false
}

func main() {
	matrix := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	target := 19
	res := searchMatrix(matrix, target)
	fmt.Println(res)
}
