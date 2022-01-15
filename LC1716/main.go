package main

func totalMoney(n int) (ans int) {
	week, day := 1, 1
	for i := 0; i < n; i++ {
		ans += week + day - 1
		day++
		if day == 8 {
			day = 1
			week++
		}
	}
	return
}
