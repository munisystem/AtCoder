package main

import "fmt"

func main() {
	var n, card int
	fmt.Scan(&n)
	cards := make([]int, n, n)
	for i := range cards {
		fmt.Scan(&card)
		cards[i] = card
	}
	fmt.Println(game(cards))
}

func game(cards []int) int {
	alice, bob := 0, 0
	t := 0
	for len(cards) != 0 {
		max, i := 0, 0
		for j := 0; j < len(cards); j++ {
			if cards[j] > max {
				max = cards[j]
				i = j
			}
		}
		if t%2 == 0 {
			alice = alice + max
		} else {
			bob = bob + max
		}
		t++
		n := []int{}
		h, t := cards[0:i], cards[i+1:]
		n = append(n, h...)
		n = append(n, t...)
		cards = n
	}
	return alice - bob
}
