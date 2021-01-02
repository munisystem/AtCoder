package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	suma, sumb := 0, 0
	for {
		suma = suma + a%10
		a = a / 10
		if a == 0 {
			break
		}
	}
	for {
		sumb = sumb + b%10
		b = b / 10
		if b == 0 {
			break
		}
	}
	fmt.Println(int(math.Max(float64(suma), float64(sumb))))
}
