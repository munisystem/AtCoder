package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	s := "No"
	if x > y {
		if x < y+3 {
			s = "Yes"
		}
	} else {
		if y < x+3 {
			s = "Yes"
		}
	}
	fmt.Println(s)
}
