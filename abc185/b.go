package main

import (
	"fmt"
	"math"
)

func main() {
	var n, m, t int
	fmt.Scan(&n, &m, &t)
	as, bs := make([]int, m, m), make([]int, m, m)
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		as[i] = a
		bs[i] = b
	}

	amount := n
	l := 0
	for i := 0; i < m; i++ {
		amount = amount - (as[i] - l)
		if amount <= 0 {
			fmt.Println("No")
			return
		}
		amount = int(math.Min(float64(amount+bs[i]-as[i]), float64(n)))
		l = bs[i]
	}
	if amount-(t-l) <= 0 {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
