package main

func sumOfUnique(nums []int) (ans int) {
	cnt := map[int]int{}
	for _, num := range nums {
		cnt[num]++
	}
	for num, c := range cnt {
		if c == 1 {
			ans += num
		}
	}
	return
}
