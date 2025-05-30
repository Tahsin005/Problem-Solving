package main

import (
	"fmt"
)

func solve() {
	var n, k int
	fmt.Scan(&n, &k)
	var s string
	fmt.Scan(&s)
	n0 := 0
	n1 := 0
	for i := 0; i < n; i++ {
		if s[i] == '1' {
			n0 += 1
		} else {
			n1 += 1
		}
	}
	mn := (max(n0, n1) - min(n0, n1)) / 2
	mx := n / 2 - n0 % 2
	if mn <= k && k <= mx && (k - mn) % 2 == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		solve()
	}
}
