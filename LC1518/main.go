package main

func numWaterBottles(numBottles int, numExchange int) (ans int) {
	for numBottles >= numExchange {
		ans += numBottles - numBottles%numExchange
		numBottles = numBottles/numExchange + numBottles%numExchange
	}
	return ans + numBottles
}
