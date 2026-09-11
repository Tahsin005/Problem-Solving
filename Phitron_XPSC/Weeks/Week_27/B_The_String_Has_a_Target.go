package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n int
		var s string
		fmt.Scan(&n, &s)
		p := 0
		for i := 1; i < n; i++ {
			if s[i] <= s[p] {
				p = i
			}
		}
		fmt.Println(string(s[p]) + s[:p] + s[p+1:])
	}
}
