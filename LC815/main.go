package main

import "math"

func numBusesToDestination(routes [][]int, source, target int) int {
	if source == target {
		return 0
	}

	n := len(routes)
	edge := make([][]bool, n)
	for i := range edge {
		edge[i] = make([]bool, n)
	}
	rec := map[int][]int{}
	for i, route := range routes {
		for _, site := range route {
			for _, j := range rec[site] {
				edge[i][j] = true
				edge[j][i] = true
			}
			rec[site] = append(rec[site], i)
		}
	}

	dis := make([]int, n)
	for i := range dis {
		dis[i] = -1
	}
	q := []int{}
	for _, site := range rec[source] {
		dis[site] = 1
		q = append(q, site)
	}
	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		for y, b := range edge[x] {
			if b && dis[y] == -1 {
				dis[y] = dis[x] + 1
				q = append(q, y)
			}
		}
	}

	ans := math.MaxInt32
	for _, site := range rec[target] {
		if dis[site] != -1 && dis[site] < ans {
			ans = dis[site]
		}
	}
	if ans < math.MaxInt32 {
		return ans
	}
	return -1
}
