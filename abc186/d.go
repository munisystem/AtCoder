package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	nums := make([]int, n, n)
	for i := 0; i < len(nums); i++ {
		var num int
		fmt.Scan(&num)
		nums[i] = num
	}

	sort.Ints(nums)
	s := 0
	for i := 0; i < n; i++ {
		s = s + nums[i]*(i*2-n+1)
	}
	fmt.Println(s)
}
