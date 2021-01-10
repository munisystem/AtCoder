package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var n, c int
	fmt.Scan(&n, &c)
	events := [][]int{}
	for i := 0; i < n; i++ {
		var a, b, c int
		fmt.Scan(&a, &b, &c)
		events = append(events, append([]int{}, a-1, c))
		events = append(events, append([]int{}, b, -c))
	}
	sort.Slice(events, func(i, j int) bool { return events[i][0] < events[j][0] })

	r := 0
	fee := 0
	t := 0
	for i := 0; i < len(events); i++ {
		r += int(math.Min(float64(c), float64(fee))) * (events[i][0] - t)
		t = events[i][0]
		fee += events[i][1]
	}
	fmt.Println(r)
}
