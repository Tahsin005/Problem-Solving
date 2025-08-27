package main

import (
	"fmt"
	"sort"
)

func main() {
	var tt int
	fmt.Scan(&tt)
	for tt > 0 {
		tt--
		var n int
		fmt.Scan(&n)

		gears := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&gears[i])
		}

		sort.Ints(gears)
		gotResult := false
		for i := 1; i < n; i++ {
			if gears[i] == gears[i - 1] {
				fmt.Println("YES")
				gotResult = true
				break
			}
		}
		if !gotResult {
			fmt.Println("NO")
		}
	}
}
