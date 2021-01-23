package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	oranges := make([]int, n, n)
	for i := 0; i < n; i++ {
		var o int
		fmt.Scan(&o)
		oranges[i] = o
	}
	ans := 0
	for i := 0; i < n; i++ {
		min := math.MaxInt32
		for j := i; j < n; j++ {
			min = int(math.Min(float64(min), float64(oranges[j])))
			ans = int(math.Max(float64(ans), float64(min*(j-i+1))))
		}
	}
	fmt.Println(ans)
}
