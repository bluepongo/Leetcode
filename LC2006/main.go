package main

func countKDifference(nums []int, k int) (ans int) {
	for j, x := range nums {
		for _, y := range nums[:j] {
			if abs(x-y) == k {
				ans++
			}
		}
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
