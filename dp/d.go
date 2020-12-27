package main

import (
	"fmt"
	"math"
)

func main() {
	var N, W int
	fmt.Scan(&N, &W)
	weights, values := make([]int, N, N), make([]int, N, N)
	for i := 0; i < N; i++ {
		var w, v int
		fmt.Scan(&w, &v)
		weights[i] = w
		values[i] = v
	}
	dp := make([][110000]int, N+1, N+1)
	for i := 0; i < len(dp[0]); i++ {
		dp[0][i] = 0
	}
	for i := 1; i < N+1; i++ {
		for j := 0; j <= W; j++ {
			if j >= weights[i-1] {
				dp[i][j] = int(math.Max(
					float64(dp[i-1][j-weights[i-1]]+values[i-1]),
					float64(dp[i-1][j]),
				))
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	fmt.Println(dp[N][W])
}
