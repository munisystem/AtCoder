package main

import "fmt"

func main() {
	var n, s, d int
	fmt.Scan(&n, &s, &d)
	magics := make([][2]int, n, n)
	for i := 0; i < n; i++ {
		var x, y int
		fmt.Scan(&x, &y)
		magics[i] = [2]int{x, y}
	}
	for i := 0; i < n; i++ {
		x, y := magics[i][0], magics[i][1]
		if x < s && y > d {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}
