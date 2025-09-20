package main

import (
	"fmt"
)

func solve() {
	var n, m int
	fmt.Scan(&n, &m)

	ans, lastT, side := 0, 0, 0
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scan(&a, &b)

		d := a - lastT
		if d % 2 == (b + side) % 2 {
			ans += d
		} else {
			ans += d - 1
		}

		lastT = a
		side = b
	}

	ans += m - lastT
	fmt.Println(ans)
}

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		solve()
	}
}
