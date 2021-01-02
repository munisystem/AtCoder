package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	if m == 0 {
		fmt.Println(1)
		return
	}
	squares := make([]int, m, m)
	for i := 0; i < m; i++ {
		var s int
		fmt.Scan(&s)
		squares[i] = s
	}
	sort.Ints(squares)
	if n == m {
		fmt.Println(0)
		return
	}

	last := 0
	k := math.MaxInt32
	t := []int{}
	for i := 0; i < len(squares); i++ {
		r := squares[i] - last - 1
		last = squares[i]
		if r == 0 {
			continue
		}
		t = append(t, r)
		k = int(math.Min(float64(k), float64(r)))
	}
	if last < n {
		t = append(t, n-last)
	}
	r := 0
	for i := 0; i < len(t); i++ {
		r = r + (t[i] / k)
		if t[i]%k != 0 {
			r = r + 1
		}
	}
	fmt.Println(r)
}
