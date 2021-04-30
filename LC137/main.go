package main

import "fmt"

func singleNumber(nums []int) int {
	freq := map[int]int{}
	for _, v := range nums {
		freq[v]++
	}
	for num, occ := range freq {
		if occ == 1 {
			return num
		}
	}
	return 0 // 不会发生，数据保证有一个元素仅出现一次
}

func main() {
	nums := []int{4, 6, 6, 6, 8, 8, 8}
	res := singleNumber(nums)
	fmt.Println(res)
}
