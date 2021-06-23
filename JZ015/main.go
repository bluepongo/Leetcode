package main

import "fmt"

func hammingWeight(num uint32) (ones int) {
	for i := 0; i < 32; i++ {
		if 1<<i&num > 0 {
			ones++
		}
	}
	return
}

func main() {
	var num uint32
	num = 00000000000000000000000000001011
	res := hammingWeight(num)
	fmt.Println(res)
}
