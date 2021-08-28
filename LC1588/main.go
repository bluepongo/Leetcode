package main

func sumOddLengthSubarrays(arr []int) int {
	var a []int
	for i := 1; i <= len(arr); i++ {
		if i%2 == 0 {
			continue
		}
		a = append(a, i)
	}
	var b [][]int
	for i := 0; i < len(arr); i++ {
		for _, l := range a {
			if l+i < len(arr) {
				b = append(b, arr[i:l+i])
			} else {
				if len(arr[i:])%2 != 0 {
					b = append(b, arr[i:])
					break
				}
			}
		}
	}
	var count = 0
	for i := range b {
		for l := range b[i] {
			count = count + b[i][l]
		}
	}
	return count
}
