package main

func maxSumOfOneSubarray(nums []int, k int) (ans []int) {
	var sum1, maxSum1 int
	for i, num := range nums {
		sum1 += num
		if i >= k-1 {
			if sum1 > maxSum1 {
				maxSum1 = sum1
				ans = []int{i - k + 1}
			}
			sum1 -= nums[i-k+1]
		}
	}
	return
}
