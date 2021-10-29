package main

import (
	"sort"
	"strconv"
)

func isPowerOfTwo(n int) bool {
	return n&(n-1) == 0
}

func reorderedPowerOf2(n int) bool {
	nums := []byte(strconv.Itoa(n))
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })

	m := len(nums)
	vis := make([]bool, m)
	var backtrack func(int, int) bool
	backtrack = func(idx, num int) bool {
		if idx == m {
			return isPowerOfTwo(num)
		}
		for i, ch := range nums {
			// 不能有前导零
			if num == 0 && ch == '0' || vis[i] || i > 0 && !vis[i-1] && ch == nums[i-1] {
				continue
			}
			vis[i] = true
			if backtrack(idx+1, num*10+int(ch-'0')) {
				return true
			}
			vis[i] = false
		}
		return false
	}
	return backtrack(0, 0)
}
