package main

import "unicode"

func reverseOnlyLetters(s string) string {
	ans := []byte(s)
	left, right := 0, len(s)-1
	for {
		for left < right && !unicode.IsLetter(rune(s[left])) { // 判断左边是否扫描到字母
			left++
		}
		for right > left && !unicode.IsLetter(rune(s[right])) { // 判断右边是否扫描到字母
			right--
		}
		if left >= right {
			break
		}
		ans[left], ans[right] = ans[right], ans[left]
		left++
		right--
	}
	return string(ans)
}
