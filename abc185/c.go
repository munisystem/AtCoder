package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	s := 1
	for i := 1; i <= 11; i++ {
		s = s * (n - i)
		s = s / i
	}
	fmt.Println(s)
}
