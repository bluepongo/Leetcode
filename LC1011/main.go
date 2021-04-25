package main

import (
	"fmt"
	"sort"
)

func shipWithinDays(weights []int, D int) int {
	// 确定二分查找左右边界
	left, right := 0, 0
	for _, w := range weights {
		if w > left {
			left = w
		}
		right += w
	}
	// Search()方法会使用“二分查找”算法来搜索某指定切片[0:n]，并返回能够使f(i)=true的最小的i（0<=i<n）
	// func Search(n int, f func(int) bool) int
	return left + sort.Search(right-left, func(x int) bool {
		x += left
		day := 1 // 需要运送的天数
		sum := 0 // 当前这一天已经运送的包裹重量之和
		for _, w := range weights {
			if sum+w > x {
				day++
				sum = 0
			}
			sum += w
		}
		return day <= D
	})
}

func main() {
	weights := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	D := 5
	res := shipWithinDays(weights, D)
	fmt.Println(res)
}
