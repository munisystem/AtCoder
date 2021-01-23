package main

import "fmt"

func main() {
	var n int
	var x int
	fmt.Scan(&n, &x)
	x = x * 100
	sakes := make([][]int, n, n)
	for i := 0; i < n; i++ {
		var v, p int
		fmt.Scan(&v, &p)
		sakes[i] = []int{v, p}
	}
	for i := 0; i < n; i++ {
		x -= sakes[i][0] * sakes[i][1]
		if x < 0 {
			fmt.Println(i + 1)
			return
		}
	}
	fmt.Println(-1)
}
