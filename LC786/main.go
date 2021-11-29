package main

import "sort"

func kthSmallestPrimeFraction(arr []int, k int) []int {
	n := len(arr)
	type pair struct{ x, y int }
	frac := make([]pair, 0, n*(n-1)/2)
	for i, x := range arr {
		for _, y := range arr[i+1:] {
			frac = append(frac, pair{x, y})
		}
	}
	sort.Slice(frac, func(i, j int) bool {
		a, b := frac[i], frac[j]
		return a.x*b.y < a.y*b.x
	})
	return []int{frac[k-1].x, frac[k-1].y}
}
