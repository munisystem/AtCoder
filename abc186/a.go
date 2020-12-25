package main

import "fmt"

func main() {
	var n, w int
	fmt.Scan(&n, &w)

	c := 0
	for ; n >= w; n = n - w {
		c++
	}
	fmt.Println(c)
}
