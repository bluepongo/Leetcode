package main

import "fmt"

// *** 爬楼梯问题 ***
// dp[i]定义：一个人跳台阶，每次可以选择跳num阶(num in nums),
// 他要跳到第i级台阶总共有多少种跳法。
// 显然，跳到第i级台阶的方法数为跳到 dp[i-num] for num in nums的方法数之和，
// 因为他只要跳到第i-num级，再一步跳num级，就可以到第i级了。

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
