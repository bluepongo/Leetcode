package main

func checkValidString(s string) bool {
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	for i, ch := range s {
		if ch == '*' {
			dp[i][i] = true
		}
	}

	for i := 1; i < n; i++ {
		c1, c2 := s[i-1], s[i]
		dp[i-1][i] = (c1 == '(' || c1 == '*') && (c2 == ')' || c2 == '*')
	}

	for i := n - 3; i >= 0; i-- {
		c1 := s[i]
		for j := i + 2; j < n; j++ {
			c2 := s[j]
			if (c1 == '(' || c1 == '*') && (c2 == ')' || c2 == '*') {
				dp[i][j] = dp[i+1][j-1]
			}
			for k := i; k < j && !dp[i][j]; k++ {
				dp[i][j] = dp[i][k] && dp[k+1][j]
			}
		}
	}

	return dp[0][n-1]
}
