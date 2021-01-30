package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	conds := make([][2]int, m, m)
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		conds[i] = [2]int{a, b}
	}
	var k int
	fmt.Scan(&k)
	people := make([][2]int, k, k)
	for i := 0; i < k; i++ {
		var c, d int
		fmt.Scan(&c, &d)
		people[i] = [2]int{c, d}
	}

	ans := 0
	for i := 0; i < 1<<k; i++ {
		balls := make([]bool, n+1, n+1)
		for j := 0; j < k; j++ {
			if (i>>j)&1 == 0 {
				balls[people[j][0]] = true
			} else {
				balls[people[j][1]] = true
			}
		}
		sum := 0
		for j := 0; j < m; j++ {
			if balls[conds[j][0]] && balls[conds[j][1]] {
				sum++
			}
		}
		if ans < sum {
			ans = sum
		}
	}
	fmt.Println(ans)
}
