package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	var a []int = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	sort.Ints(a)
	var need int = 0
	for i := 0; i < n; i += 2 {
		need += a[i + 1] - a[i]
	}
	fmt.Println(need)
}
