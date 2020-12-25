package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(oddOrEven(a, b))
}

func oddOrEven(a, b int) string {
	if a*b%2 == 0 {
		return "Even"
	}
	return "Odd"
}
