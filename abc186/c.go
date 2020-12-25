package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	c := 0
	for i := 1; i <= n; i++ {
		b := false
		i8 := fmt.Sprintf("%o", i)
		for j := 0; j < len(i8); j++ {
			if i8[j:j+1] == "7" {
				b = true
				break
			}
		}
		i10 := fmt.Sprintf("%d", i)
		for j := 0; j < len(i10); j++ {
			if i10[j:j+1] == "7" {
				b = true
				break
			}
		}
		if !b {
			c++
		}
	}
	fmt.Println(c)
}
