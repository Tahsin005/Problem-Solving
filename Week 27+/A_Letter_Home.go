package main

import (
	"fmt"
	"slices"
)

func dist(a int, b int) int {
	if a <= b {
		return b - a
	} else {
		return a - b
	}
}

func MIN(a int, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func main() {
	var t int
	fmt.Scan(&t)

	for t > 0 {
		var n, s int
		fmt.Scan(&n, &s)

		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&arr[i])
		}

		max := slices.Max(arr)
		min := slices.Min(arr)

		distance := MIN(dist(s, max) + dist(max, min), dist(s, min) + dist(min, max))

		fmt.Println(distance)

		t--
	}
}