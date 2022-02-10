package main

import "strconv"

func simplifiedFractions(n int) (ans []string) {
	for denominator := 2; denominator <= n; denominator++ {
		for numerator := 1; numerator < denominator; numerator++ {
			if gcd(numerator, denominator) == 1 {
				ans = append(ans, strconv.Itoa(numerator)+"/"+strconv.Itoa(denominator))
			}
		}
	}
	return
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
