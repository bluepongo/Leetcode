package main

import "strconv"

type node struct {
	minVal, maxVal float64
	minStr, maxStr string
}

func optimalDivision(nums []int) string {
	n := len(nums)
	dp := make([][]node, n)
	for i := range dp {
		dp[i] = make([]node, n)
		for j := range dp[i] {
			dp[i][j].minVal = 1e4
		}
	}

	for i, num := range nums {
		dp[i][i].minVal = float64(num)
		dp[i][i].maxVal = float64(num)
		dp[i][i].minStr = strconv.Itoa(num)
		dp[i][i].maxStr = strconv.Itoa(num)
	}
	for i := 1; i < n; i++ {
		for j := 0; j+i < n; j++ {
			for k := j; k < j+i; k++ {
				if dp[j][j+i].maxVal < dp[j][k].maxVal/dp[k+1][j+i].minVal {
					dp[j][j+i].maxVal = dp[j][k].maxVal / dp[k+1][j+i].minVal
					if k+1 == j+i {
						dp[j][j+i].maxStr = dp[j][k].maxStr + "/" + dp[k+1][j+i].minStr
					} else {
						dp[j][j+i].maxStr = dp[j][k].maxStr + "/(" + dp[k+1][j+i].minStr + ")"
					}
				}
				if dp[j][j+i].minVal > dp[j][k].minVal/dp[k+1][j+i].maxVal {
					dp[j][j+i].minVal = dp[j][k].minVal / dp[k+1][j+i].maxVal
					if k+1 == j+i {
						dp[j][j+i].minStr = dp[j][k].minStr + "/" + dp[k+1][j+i].maxStr
					} else {
						dp[j][j+i].minStr = dp[j][k].minStr + "/(" + dp[k+1][j+i].maxStr + ")"
					}
				}
			}
		}
	}
	return dp[0][n-1].maxStr
}
