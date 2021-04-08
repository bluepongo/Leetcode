package main

import "fmt"

func findMin(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		pivot := low + (high-low)/2
		if nums[pivot] < nums[high] {
			high = pivot
		} else {
			low = pivot + 1
		}
	}
	return nums[low]
}

func main() {
	nums := []int{2, 5, 6, 0, 0, 1, 2}
	res := findMin(nums)
	fmt.Println(res)
}
