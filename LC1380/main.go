package main

func luckyNumbers(matrix [][]int) (ans []int) {
	for _, row := range matrix {
	next:
		for j, x := range row {
			for _, y := range row {
				if y < x {
					continue next
				}
			}
			for _, r := range matrix {
				if r[j] > x {
					continue next
				}
			}
			ans = append(ans, x)
		}
	}
	return
}
