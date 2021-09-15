package main

func findPeakElement(nums []int) (idx int) {
	for i, v := range nums {
		if v > nums[idx] {
			idx = i
		}
	}
	return
}
