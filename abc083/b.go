package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)
	fmt.Println(sum(n, a, b))
}

func sum(n, a, b int) int {
	s := 0
	for i := 1; i <= n; i++ {
		ss := 0
		for j := i; j > 0; j = j / 10 {
			ss = ss + j%10
		}
		if ss >= a && ss <= b {
			s = s + i
		}
	}
	return s
}
