package main

import (
	"fmt"
	"math"
)

func main() {
	r := math.MaxInt32
	for i := 0; i < 4; i++ {
		var n int
		fmt.Scan(&n)
		r = int(math.Min(float64(r), float64(n)))
	}
	fmt.Println(r)
}
