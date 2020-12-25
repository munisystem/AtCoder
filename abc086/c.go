package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	points := make([][]int, n, n)
	for i := range points {
		var t, x, y int
		fmt.Scan(&t, &x, &y)
		points[i] = []int{t, x, y}
	}
	fmt.Println(travel(points))
}

func travel(points [][]int) string {
	pt, px, py := 0, 0, 0
	for i := 0; i < len(points); i++ {
		t, x, y := points[i][0], points[i][1], points[i][2]
		dt := t - pt
		dx, dy := 0, 0
		if px >= x {
			dx = px - x
		} else {
			dx = x - px
		}
		if py >= y {
			dy = py - y
		} else {
			dy = y - py
		}
		if dt < dx+dy || dt%2 != (dx+dy)%2 {
			return "No"
		}
		pt, px, py = t, x, y
	}
	return "Yes"
}
