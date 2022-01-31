package main

func numberOfSteps(num int) (ans int) {
	for ; num > 0; num >>= 1 {
		ans += num & 1
		if num > 1 {
			ans++
		}
	}
	return
}
