package main

import "fmt"

func rotate(mat [][]int) {
	n := len(mat)
	for i := 0; i < n/2; i++ {
		for j := 0; j < (n+1)/2; j++ {
			mat[i][j], mat[j][n-1-i], mat[n-1-i][n-1-j], mat[n-1-j][i] =
				mat[n-1-j][i], mat[i][j], mat[j][n-1-i], mat[n-1-i][n-1-j]
		}
	}
}

func main() {
	matirx := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(matirx)
	for _, m := range matirx {
		fmt.Println(m)
	}
	fmt.Println(matirx)
}
