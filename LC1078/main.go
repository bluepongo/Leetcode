package main

import "strings"

func findOcurrences(text, first, second string) (ans []string) {
	words := strings.Split(text, " ")
	for i := 2; i < len(words); i++ {
		if words[i-2] == first && words[i-1] == second {
			ans = append(ans, words[i])
		}
	}
	return
}
