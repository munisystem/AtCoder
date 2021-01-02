package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	xs, ys := make([]int, n, n), make([]int, n, n)
	for i := 0; i < n; i++ {
		var x, y int
		fmt.Scan(&x, &y)
		xs[i] = x
		ys[i] = y
	}
	r := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			t := float64(ys[j]-ys[i]) / float64(xs[j]-xs[i])
			if t >= -1 && t <= 1 {
				r++
			}
		}
	}
	fmt.Println(r)
}
