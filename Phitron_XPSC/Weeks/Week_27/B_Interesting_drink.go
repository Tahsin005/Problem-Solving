package main

import (
	"fmt"
	"sort"
)

func main () {
	var n, q int
	fmt.Scan(&n)

	x := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x[i])
	}

	fmt.Scan(&q)
	m := make([]int, q)
	for i := 0; i < q; i++ {
		fmt.Scan(&m[i])
	}

	sort.Ints(x)
	for _, money := range m {
		val := sort.Search(len(x), func(i int) bool {
			return x[i] > money
		})
		fmt.Println(val)
	}
}