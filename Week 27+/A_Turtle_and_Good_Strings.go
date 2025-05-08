package main

import (
	"fmt"
)

func solve() {
	var n int
	fmt.Scan(&n)
	var s string
	fmt.Scan(&s)
	if len(s) >= 2 && s[0] != s[len(s) - 1] {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
func main() {
	var t int
	fmt.Scan(&t)
	for ; t > 0; t-- {
		solve()
	}
}
