package main

func isCovered(ranges [][]int, left int, right int) bool {
	for left <= right {
		have := false
		for _, v := range ranges {
			if left >= v[0] && left <= v[1] {
				have = true
				break
			}
		}
		if have == false {
			return false
		}
		left++
	}
	return true
}
