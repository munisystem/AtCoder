package main

import "fmt"

func main() {
	var a string
	fmt.Scan(&a)
	if a[0] == a[1] && a[1] == a[2] {
		fmt.Println("Won")
	} else {
		fmt.Println("Lost")
	}
}
