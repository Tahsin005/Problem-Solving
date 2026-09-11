package main

import (
	"fmt"
)

func ok(x, y int) bool {
	if x < y {
		x, y = y, x
	}
	return x <= 2 * (y + 1)
}

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var a, b, c, d int
		fmt.Scan(&a, &b, &c, &d)

		firstHalf := ok(a, b)
		secondHalf := ok(c - a, d - b)

		if firstHalf && secondHalf {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
