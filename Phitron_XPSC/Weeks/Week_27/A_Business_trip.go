package main

import (
	"fmt"
	"sort"
)

func main() {
	var k int
	fmt.Scan(&k)
	if k == 0 {
		fmt.Println(0)
		return
	}
	a := make([]int, 12)
	for i := range a {
		fmt.Scan(&a[i])
	}
	sort.Ints(a)
	for i := 11; i >= 0; i-- {
		if k <= a[i] {
			fmt.Println(12 - i)
			return
		}
		k -= a[i]
	}
	fmt.Println(-1)
}
