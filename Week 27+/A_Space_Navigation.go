package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func solve() {
	f := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	N, x, y, S := 0, 0, 0, ""
	fmt.Fscan(f, &N)
	for i := 0; i < N; i++ {
		fmt.Fscan(f, &x, &y, &S)
		u, d, r, l := 0, 0, 0, 0
		U, D, R, L := strings.Count(S, "U"), strings.Count(S, "D"), strings.Count(S, "R"), strings.Count(S, "L")
		if x > 0 {
			r = x
		}
		if x < 0 {
			l = x
		}
		if y > 0 {
			u = y
		}
		if y < 0 {
			d = y
		}
		if u > 0 && U < u || d < 0 && D < -d || r > 0 && R < r || l < 0 && L < -l {
			fmt.Fprintln(w, "NO")
		} else {
			fmt.Fprintln(w, "YES")
		}
	}
	w.Flush()
}

func main() {
	solve()
}
