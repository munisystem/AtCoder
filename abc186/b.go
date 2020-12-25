package main

import (
	"fmt"
	"sort"
)

func main() {
	var h, w int
	fmt.Scan(&h, &w)
	nums := make([]int, h*w, h*w)
	for i := 0; i < len(nums); i++ {
		var n int
		fmt.Scan(&n)
		nums[i] = n
	}
	sort.Ints(nums)
	sum, b := 0, nums[0]
	for i := 1; i < len(nums); i++ {
		sum = sum + (nums[i] - b)
	}

	fmt.Println(sum)
}
