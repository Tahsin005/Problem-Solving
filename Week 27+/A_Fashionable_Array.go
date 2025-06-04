package main

import (
	"fmt"
	"sort"
)

func solve() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n, n)
	for i := range n {
		fmt.Scan(&a[i])
	}
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	res := n
	for i := 0; i < n; i++ {
		if a[i] % 2 == a[n - 1] % 2 {
			res = i
			break
		}
	}
	for i := n - 1; i >= 0; i-- {
		if a[0] % 2 == a[i] % 2 {
			res = min(n - i - 1, res)
			break
		}
	}
	fmt.Println(res)
}

func main() {
	var t int
	fmt.Scan(&t)
	for range t {
		solve()
	}
}
