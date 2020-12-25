package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	b := daydream(s)
	if b {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

var words = []string{"dream", "dreamer", "eraser", "erase"}

func daydream(s string) bool {
	if len(s) == 0 {
		return true
	}
	b := false
	for _, word := range words {
		if len(s) >= len(word) {
			if p := s[0:len(word)]; p == word {
				b = b || daydream(s[len(word):])
			}
		}
	}
	return b
}
