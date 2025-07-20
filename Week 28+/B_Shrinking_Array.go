package main

import (
	"fmt"
)

func intAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func solve(a []int, n int) int {
	for i := 0; i < n - 1; i++ {
		if intAbs(a[i + 1] - a[i]) <= 1 {
			return 0
		}
	}

	for i := 1; i < n - 1; i++ {
		if a[i - 1] <= a[i] && a[i + 1] <= a[i] {
			return 1
		}
		if a[i - 1] >= a[i] && a[i + 1] >= a[i] {
			return 1
		}
	}

	return -1
}

func main() {
	var t int

	for fmt.Scan(&t); t > 0; t-- {
		var n int
		fmt.Scan(&n)

		a := make([]int, n)

		for i := 0; i < n; i++ {
			fmt.Scan(&a[i])
		}

		fmt.Println(solve(a, n))
	}
}
