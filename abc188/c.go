package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	an := int(math.Pow(2, float64(n)))
	queue := [][]int{}
	for i := 1; i <= an; i++ {
		var a int
		fmt.Scan(&a)
		queue = append(queue, append([]int{}, i, a))
	}
	for i := 1; i < n; i++ {
		l := len(queue) / 2
		for j := 0; j < l; j++ {
			a1, a2 := queue[0], queue[1]
			queue = queue[2:]
			if a1[1] > a2[1] {
				queue = append(queue, a1)
			} else {
				queue = append(queue, a2)
			}
		}
	}
	a1, a2 := queue[0], queue[1]
	if a1[1] < a2[1] {
		fmt.Println(a1[0])
	} else {
		fmt.Println(a2[0])
	}
}
