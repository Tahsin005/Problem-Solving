package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, m int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	fmt.Scan(&m)
	b := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&b[i])
	}
	sort.Ints(a)
	sort.Ints(b)
	res := 0
	for i, j := 0, 0; i < n && j < m; {
		switch {
		case a[i] + 1 < b[j]:
			i++
		case b[j] + 1 < a[i]:
			j++
		default:
			res++
			i++
			j++
		}
	}
	fmt.Println(res)
}
