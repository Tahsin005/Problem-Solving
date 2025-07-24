package main

import (
	"fmt"
	"sort"
)

func main () {
	var d1, d2, d3 int
	fmt.Scan(&d1, &d2, &d3)

	x := []int{d1 + d2 + d3, d1 * 2 + d2 * 2, d1 * 2 + d3 * 2, d2 * 2 + d3 * 2}
	sort.Ints(x)
	fmt.Println(x[0])
}