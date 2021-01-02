package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	ss := make([]string, n, n)
	for i := 0; i < n; i++ {
		var s string
		fmt.Scan(&s)
		ss[i] = s
	}
	m := make(map[string]interface{})
	for i := 0; i < n; i++ {
		m[ss[i]] = nil
	}
	for i := 0; i < n; i++ {
		s := ss[i]
		if _, ok := m["!"+s]; ok {
			fmt.Println(s)
			return
		}
	}
	fmt.Println("satisfiable")
}
