package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var tt int
	fmt.Fscan(in, &tt)

	for ; tt > 0; tt-- {
		var n, k int
		fmt.Fscan(in, &n, &k)

		type pair struct {
			l, r, real int
		}

		a := make([]pair, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &a[i].l, &a[i].r, &a[i].real)
		}

		sort.Slice(a, func(i, j int) bool {
			return a[i].real < a[j].real
		})

		for _, p := range a {
			if p.l <= k && p.r >= k {
				k = max(k, p.real)
			}
		}

		fmt.Println(k)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
