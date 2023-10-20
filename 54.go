package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	dx := []int{0, 1, 0, -1}
	dy := []int{1, 0, -1, 0}
	x, y, k := 0, -1, 0
	m, n := len(matrix), len(matrix[0])
	cnt := m * n
	var res []int
	m--
	for cnt > 0 {
		for i := 0; i < n; i++ {
			x, y = x+dx[k], y+dy[k]
			res = append(res, matrix[x][y])
		}
		cnt -= n
		n--
		k = (k + 1) % 4
		for i := 0; i < m; i++ {
			x, y = x+dx[k], y+dy[k]
			res = append(res, matrix[x][y])
		}
		cnt -= m
		m--
		k = (k + 1) % 4
	}
	return res
}

func spiralOrder0(matrix [][]int) []int {
	dx := []int{0, 1, 0, -1}
	dy := []int{1, 0, -1, 0}
	i, j, k := 0, -1, 0
	m, n := len(matrix), len(matrix[0])
	var res []int
	for c := 0; c < m*n; c++ {
		x, y := i+dx[k], j+dy[k]
		if x < 0 || x >= m || y < 0 || y >= n || matrix[x][y] == -1 {
			k = (k + 1) % 4
			x, y = i+dx[k], j+dy[k]
		}
		res = append(res, matrix[x][y])
		matrix[x][y] = -1
		i, j = x, y
	}
	return res
}

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	res := spiralOrder(matrix)
	fmt.Println(res)
}
