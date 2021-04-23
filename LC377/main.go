package main

import "fmt"

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	// 外层循环遍历状态
	for i := 1; i <= target; i++ {
		// 里层循环遍历选择
		for _, num := range nums {
			if i-num >= 0 {
				// 所有可能的选择之和
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}

func main() {
	nums := []int{1, 2, 3}
	target := 4
	res := combinationSum4(nums, target)
	fmt.Println(res)
}
