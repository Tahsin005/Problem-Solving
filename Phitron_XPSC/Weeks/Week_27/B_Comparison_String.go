package main

import (
	"fmt"
)

func solve() {
	var n int
	fmt.Scan(&n)
	var s string
	fmt.Scan(&s)
	ml, cl := 1, 1
	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			cl++
		} else {
			if cl > ml {
				ml = cl
			}
			cl = 1
		}
	}
	if cl > ml {
		ml = cl
	}
	fmt.Println(ml + 1)
}

func main() {
	var t int
	fmt.Scan(&t)

	for t > 0 {
		solve()
		t--
	}
}
