package main

import "fmt"

func main() {
	var fh, h, fi, p int
	fmt.Scan(&fh, &h, &fi, &p)
	fmt.Println(coins([]int{fh, h, fi}, p))
}

func coins(nums []int, price int) int {
	counter := 0
	for i := 0; i <= nums[0]; i++ {
		for j := 0; j <= nums[1]; j++ {
			for k := 0; k <= nums[2]; k++ {
				if ((i * 500) + (j * 100) + (k * 50)) == price {
					counter++
				}
			}
		}
	}
	return counter
}
