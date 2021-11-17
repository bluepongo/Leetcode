package main

import "strings"

// 暴力解法
// m 表示单词的平均长度，n 表示单词的个数
// 时间复杂度：O(n^2 * m)
// 空间复杂度：O(1)
func maxProduct(words []string) int {
	var ans = 0
	for i := range words {
		var word1 = words[i]
		for j := i + 1; j < len(words); j++ {
			var word2 = words[j]
			// 每个单词的 bitMask 会重复计算很多次
			if !hasSameChar(word1, word2) {
				var length = len(word1) * len(word2)
				if length > ans {
					ans = length
				}
			}
		}
	}
	return ans
}

// O(m^2)
func hasSameChar(word1 string, word2 string) bool {
	for _, c := range word1 {
		if strings.Index(word2, string(c)) >= 0 {
			return true
		}
	}
	return false
}
