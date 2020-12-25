package main

import "fmt"

func main() {
	var a string
	fmt.Scan(&a)
	fmt.Println(numOfMarbles(a))
}

func numOfMarbles(a string) int {
	counter := 0
	for i := 0; i < len(a); i++ {
		if a[i:i+1] == "1" {
			counter++
		}
	}
	return counter
}
