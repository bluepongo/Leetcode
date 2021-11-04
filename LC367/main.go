package main

import "math"

func isPerfectSquare(num int) bool {
	x := int(math.Sqrt(float64(num)))
	return x*x == num
}
