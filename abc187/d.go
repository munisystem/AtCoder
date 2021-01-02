package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	votes := make([]int, n, n)
	v := 0
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		v = v - a
		votes[i] = 2*a + b
	}
	sort.Ints(votes)
	r := 0
	for i := len(votes) - 1; i > 0; i-- {
		v = v + votes[i]
		r++
		if v > 0 {
			break
		}
	}
	fmt.Println(r)
}
