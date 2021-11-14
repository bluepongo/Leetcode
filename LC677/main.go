package main

import "strings"

type MapSum map[string]int

func Constructor() MapSum {
	return MapSum{}
}

func (m MapSum) Insert(key string, val int) {
	m[key] = val
}

func (m MapSum) Sum(prefix string) (sum int) {
	for s, v := range m {
		if strings.HasPrefix(s, prefix) {
			sum += v
		}
	}
	return
}
