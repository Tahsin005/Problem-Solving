package main

import (
	"fmt"
	"sort"
)

func main () {
	var n, q int
	fmt.Scan(&n, &q)

	p := make([]int64, n)

	for i := range p {
		fmt.Scan(&p[i])
	}

	sort.Slice(p, func(i, j int) bool {
		return p[i] < p[j]
	})

	psa := make([]int64, n)

	for i := range p {
		psa[i] = p[i]

		if i != 0 {
			psa[i] += psa[i-1]
		}
	}

	for i := 0; i < q; i++ {
		var x, y int
		fmt.Scan(&x, &y)

		l := n - x
		r := n - x + y - 1

		s := psa[r] - psa[l] + p[l]
		fmt.Println(s)
	}
}