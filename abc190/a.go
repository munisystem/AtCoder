package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if a > b || (a == b && c == 1) {
		fmt.Println("Takahashi")
	} else {
		fmt.Println("Aoki")
	}
}
