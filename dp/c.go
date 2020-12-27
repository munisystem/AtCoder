package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	d := make([][3]int, n, n)
	for i := 0; i < n; i++ {
		var a, b, c int
		fmt.Scan(&a, &b, &c)
		d[i][0] = a
		d[i][1] = b
		d[i][2] = c
	}
	dp := make([][3]int, n+1, n+1)
	dp[0][0] = 0
	dp[0][1] = 0
	dp[0][2] = 0

	for i := 1; i < n+1; i++ {
		for j := 0; j <= 2; j++ {
			switch j {
			case 0:
				dp[i][0] = d[i-1][0] + int(math.Max(float64(dp[i-1][1]), float64(dp[i-1][2])))
			case 1:
				dp[i][1] = d[i-1][1] + int(math.Max(float64(dp[i-1][0]), float64(dp[i-1][2])))
			case 2:
				dp[i][2] = d[i-1][2] + int(math.Max(float64(dp[i-1][0]), float64(dp[i-1][1])))
			}
		}
	}
	max := 0
	for i := 0; i < len(dp[n]); i++ {
		if max < dp[n][i] {
			max = dp[n][i]
		}
	}
	fmt.Println(max)
}
