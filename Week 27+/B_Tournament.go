package main

import (
	"fmt"
	"sort"
)

func main () {
	var t int
	fmt.Scan(&t)
	for range t {
		var n, j, k int
		fmt.Scan(&n, &j, &k)
		a := make([]int, n)
		for i := range n {
			fmt.Scan(&a[i])
		}
		val := a[j - 1]
		sort.Ints(a)
		if k >= 2 || a[n - 1] == val {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}