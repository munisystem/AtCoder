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
	fmt.Println(tower(nums))
}

func tower(nums []int) int {
	f := 0
	l := 0
	for len(nums) > 0 {
		m, i := 0, 0
		for j := 0; j < len(nums); j++ {
			if m < nums[j] {
				m = nums[j]
				i = j
			}
		}
		if m == 0 {
			break
		}
		if !(l != 0 && l == nums[i]) {
			f++
		}
		l = nums[i]
		n := []int{}
		h, t := nums[0:i], nums[i+1:]
		n = append(n, h...)
		n = append(n, t...)
		nums = n
	}
	return f
}
