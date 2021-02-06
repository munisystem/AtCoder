package main

import (
	"fmt"
	"strings"
)

func main() {
	var n, x int
	fmt.Scan(&n, &x)
	ans := []string{}
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		if a != x {
			ans = append(ans, fmt.Sprintf("%d", a))
		}
	}
	fmt.Println(strings.Join(ans, " "))
}
