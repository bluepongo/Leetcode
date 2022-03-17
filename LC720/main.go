package main

import "sort"

func longestWord(words []string) (longest string) {
	sort.Slice(words, func(i, j int) bool {
		s, t := words[i], words[j]
		return len(s) < len(t) || len(s) == len(t) && s > t
	})

	candidates := map[string]struct{}{"": {}}
	for _, word := range words {
		if _, ok := candidates[word[:len(word)-1]]; ok {
			longest = word
			candidates[word] = struct{}{}
		}
	}
	return longest
}
