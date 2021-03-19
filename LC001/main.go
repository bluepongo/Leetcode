package main

import "fmt"

func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}

func main(){
	nums := []int{15, 16, 18, 2, 7}
	target := 9
	res := twoSum(nums, target)
	fmt.Println(res)
}
