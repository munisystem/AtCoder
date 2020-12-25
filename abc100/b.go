package main

import (
	"fmt"
)

func main() {
	var d, n int
	fmt.Scan(&d, &n)
	p := 1
	for i := 0; i < d; i++ {
		p = p * 100
	}
	if n == 100 {
		fmt.Println(p*100 + 1*p)
	} else {
		fmt.Println((p * n))
	}
}
