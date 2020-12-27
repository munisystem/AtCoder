package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	h := make([]int, n, n)
	for i := 0; i < n; i++ {
		var j int
		fmt.Scan(&j)
		h[i] = j
	}
	dp := make([]int, n, n)
	dp[0] = 0
	for i := 1; i < n; i++ {
		if i == 1 {
			dp[i] = int(math.Abs(float64(h[i] - h[i-1])))
			continue
		}
		dp[i] = int(math.Min(
			float64(dp[i-1])+math.Abs(float64(h[i]-h[i-1])),
			float64(dp[i-2])+math.Abs(float64(h[i]-h[i-2])),
		))
	}
	fmt.Println(dp[n-1])
}
