package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t)

	for t > 0 {
		t--
		var x, y, z, m int

		fmt.Scan(&x, &y, &z)
		m = max(x, y, z)

		if (m == y && m == z) ||
			(m == x && m == z) ||
			(m == x && m == y) {
			fmt.Print("YES\n")
			fmt.Printf("%d %d %d\n", m, min(x, y, z), min(x, y, z))
		} else {
			fmt.Print("NO\n")
		}
	}
}
