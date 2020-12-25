package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	nums := make([]int, n, n)
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		nums[i] = num
	}

	c := 0
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		for ; num%2 == 0; num = num / 2 {
			c++
		}
	}
	fmt.Println(c)
}
