package main

import "sort"

func searchMatrix(matrix [][]int, target int) bool {
	row := sort.Search(len(matrix), func(i int) bool { return matrix[i][0] > target }) - 1
	if row < 0 {
		return false
	}
	col := sort.SearchInts(matrix[row], target)
	return col < len(matrix[row]) && matrix[row][col] == target
}

func searchMatrix002(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	i := sort.Search(m*n, func(i int) bool { return matrix[i/n][i%n] >= target })
	return i < m*n && matrix[i/n][i%n] == target
}

func main() {
	mat := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}
	searchMatrix(mat, 11)
}
