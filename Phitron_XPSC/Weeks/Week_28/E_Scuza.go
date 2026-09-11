package main

import (
	"bufio"
	"fmt"
	"os"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func bsearch(a []int, l, r, v int) int {
	if (l == r) {
		return l
	}
	m := (l + r + 1) >> 1
	if a[m] > v {
		return bsearch(a, l, m-1, v)
	}
	return bsearch(a, m, r, v)
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var T int
	for fmt.Fscan(r, &T); T > 0; T-- {
		var n, q int
		fmt.Fscan(r, &n, &q)
		a := make([]int, n + 1)
		h := make([]int, n + 1)
		for i := range a {
			if i == 0 { continue }
			fmt.Fscan(r, &a[i])
			h[i] = h[i - 1] + a[i]
			a[i] = max(a[i], a[i - 1])
		}

		for i := 0; i < q; i++ {
			var qq int
			fmt.Fscan(r, &qq)
			fmt.Fprintf(w, "%d ", h[bsearch(a, 0, n, qq)])
		}

		fmt.Fprintln(w, "")
	}
}
