package main

import (
	"fmt"
	"math"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	h := make([]int, n, n)
	for i := 0; i < n; i++ {
		var j int
		fmt.Scan(&j)
		h[i] = j
	}
	dp := make([]int, n, n)
	for i := 0; i < n; i++ {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := 1; i < n; i++ {
		for j := 1; j <= k; j++ {
			if j > i {
				continue
			}
			dp[i] = int(math.Min(
				float64(dp[i]),
				float64(dp[i-j])+math.Abs(float64(h[i]-h[i-j])),
			))
		}
	}
	fmt.Println(dp[n-1])
}
