package main

import "fmt"

func main() {
	var n, y int
	fmt.Scan(&n, &y)
	fmt.Println(otoshidama(n, y))
}

func otoshidama(n, y int) (int, int, int) {
	for i := 0; i <= n; i++ {
		for j := 0; j <= (n - i); j++ {
			k := n - i - j
			if (i*10000)+(j*5000)+(k*1000) == y {
				return i, j, k
			}
		}
	}
	return -1, -1, -1
}
