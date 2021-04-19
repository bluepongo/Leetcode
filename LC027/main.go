package main

import "fmt"

func removeElement(nums []int, val int) int {
	left, right := 0, len(nums)
	for left < right {
		if nums[left] == val {
			nums[left] = nums[right-1]
			right--
		} else {
			left++
		}
	}
	return left
}

func main() {
	num := []int{2, 2, 4, 4}
	val := 4
	res := removeElement(num, val)
	for i := 0; i < res; i++ {
		fmt.Println(num[i])
	}
}
