package main

import "fmt"

func countTriplets(arr []int) (ans int) {
	n := len(arr)
	s := make([]int, n+1)
	for i, v := range arr {
		s[i+1] = s[i] ^ v
	}
	cnt := map[int]int{}
	total := map[int]int{}
	for k := 0; k < n; k++ {
		if m, has := cnt[s[k+1]]; has {
			ans += m*k - total[s[k+1]]
		}
		cnt[s[k]]++
		total[s[k]] += k
	}
	return
}

func main() {
	arr := []int{2, 3, 1, 6, 7}
	res := countTriplets(arr)
	fmt.Println(res)
}
