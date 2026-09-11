package main

import (
	"fmt"
	"slices"
)

func main() {
	var t int
	fmt.Scan(&t)
	for ; t > 0; t-- {
		var n int
		fmt.Scan(&n)
		m := make(map[int]int)
		for _ = range n {
			var x int
			fmt.Scan(&x)
			m[x]++
		}

		ans := 0
		var a []int
		for _, v := range m {
			a = append(a, v)
		}

		slices.Sort(a)
		for i := len(a) - 1; i >= 0; i-- {
			ans = max(ans, a[i] * (len(a) - i))
		}
		fmt.Println(ans)
	}
}
