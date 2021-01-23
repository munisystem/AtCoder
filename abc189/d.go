package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	s := make([]string, n, n)
	for i := 0; i < n; i++ {
		var b string
		fmt.Scan(&b)
		s[i] = b
	}
	m := make([][]int, n+1, n+1)
	for i := 0; i <= n; i++ {
		m[i] = make([]int, 2, 2)
	}
	m[0][0] = 1
	m[0][1] = 1
	for i := 0; i < n; i++ {
		b := s[i]
		if b == "OR" {
			m[i+1][0] = 2*m[i][0] + m[i][1]
			m[i+1][1] = m[i][1]
		} else {
			m[i+1][0] = m[i][0]
			m[i+1][1] = m[i][0] + 2*m[i][1]
		}
	}
	fmt.Println(m[n][0])
}
