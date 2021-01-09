package main

import "fmt"

func main() {
	var n, x int
	var s string
	fmt.Scan(&n, &x, &s)
	r := x
	for i := 0; i < n; i++ {
		if string(s[i]) == "o" {
			r++
			continue
		}
		if r == 0 {
			continue
		}
		r--
	}
	fmt.Println(r)
}
