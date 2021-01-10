package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	as, bs := make([]int, n, n), make([]int, n, n)
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		as[i] = a
	}
	for i := 0; i < n; i++ {
		var b int
		fmt.Scan(&b)
		bs[i] = b
	}
	r := 0
	for i := 0; i < n; i++ {
		r = r + as[i]*bs[i]
	}
	if r == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
