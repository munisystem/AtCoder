package main

import "fmt"

func main() {
	var n, num int
	fmt.Scan(&n)
	nums := make([]int, n, n)
	for i := range nums {
		fmt.Scan(&num)
		nums[i] = num
	}
	fmt.Println(shift(nums))
}

func shift(nums []int) int {
	counter := 0
	c := true
	for c {
		for i, num := range nums {
			if num%2 != 0 {
				c = false
				break
			}
			nums[i] = num / 2
		}
		if c {
			counter++
		}
	}
	return counter
}
