package main

import (
	"fmt"
	"sort"
)

type pair struct {
	a, b int
}

func main() {
	var n, res int
	var p []pair
	fmt.Scanf("%d\n", &n)
	p = make([]pair, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d\n", &p[i].a, &p[i].b)
	}
	sort.SliceStable(p, func(i, j int) bool {
		return p[i].a < p[j].a || (p[i].a == p[j].a && p[i].b < p[j].b)
	})
	for i := 0; i < n; i++ {
		if res > p[i].b {
			res = p[i].a
		} else {
			res = p[i].b
		}
	}
	fmt.Println(res)
}