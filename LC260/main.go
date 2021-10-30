package main

func singleNumber(nums []int) (ans []int) {
	freq := map[int]int{}
	for _, num := range nums {
		freq[num]++
	}
	for num, occ := range freq {
		if occ == 1 {
			ans = append(ans, num)
		}
	}
	return
}
