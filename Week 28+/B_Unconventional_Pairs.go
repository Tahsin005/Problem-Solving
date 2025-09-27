package main

import (
	"fmt"
	"sort"
)

func main() {
	var t int
	fmt.Scan(&t)

	for t > 0 {
		t--
		var n int
		fmt.Scan(&n)

		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&a[i])
		}

		sort.Ints(a)

		ans := 0
		for i := 0; i < n; i += 2 {
			ans = max(ans, a[i + 1] - a[i])
		}

		fmt.Println(ans)
	}
}