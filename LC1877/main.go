package main

import "sort"

func minPairSum(nums []int) (ans int) {
	sort.Ints(nums)
	n := len(nums)
	for i, val := range nums[:n/2] {
		ans = max(ans, val+nums[n-1-i])
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
