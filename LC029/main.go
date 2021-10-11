package main

import "math"

// 时间复杂度：O(1)
func divide(a int, b int) int {
	if a == math.MinInt32 && b == -1 {
		return math.MaxInt32
	}

	sign := 1
	if (a > 0 && b < 0) || (a < 0 && b > 0) {
		sign = -1
	}

	a = abs(a)
	b = abs(b)

	res := 0
	for i := 31; i >= 0; i-- {
		if (a>>i)-b >= 0 {
			a = a - (b << i)
			res += 1 << i
		}
	}
	return sign * res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
