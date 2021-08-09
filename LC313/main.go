package main

import "math"

func nthSuperUglyNumber(n int, primes []int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	m := len(primes)
	pointers := make([]int, m)
	for i := range pointers {
		pointers[i] = 1
	}
	for i := 2; i <= n; i++ {
		nums := make([]int, m)
		minNum := math.MaxInt64
		for j, p := range pointers {
			nums[j] = dp[p] * primes[j]
			minNum = min(minNum, nums[j])
		}
		dp[i] = minNum
		for j, num := range nums {
			if minNum == num {
				pointers[j]++
			}
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
