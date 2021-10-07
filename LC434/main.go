package main

func countSegments(s string) (ans int) {
	for i, ch := range s {
		if (i == 0 || s[i-1] == ' ') && ch != ' ' {
			ans++
		}
	}
	return
}
