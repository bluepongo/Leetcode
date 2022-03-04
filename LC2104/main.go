package main

func subArrayRanges(nums []int) (ans int64) {
	for i, num := range nums {
		minVal, maxVal := num, num
		for _, v := range nums[i+1:] {
			if v < minVal {
				minVal = v
			} else if v > maxVal {
				maxVal = v
			}
			ans += int64(maxVal - minVal)
		}
	}
	return
}
