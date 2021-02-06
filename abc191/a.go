package main

import "fmt"

func main() {
	var v, t, s, d int
	fmt.Scan(&v, &t, &s, &d)
	if v*t > d || v*s < d {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
