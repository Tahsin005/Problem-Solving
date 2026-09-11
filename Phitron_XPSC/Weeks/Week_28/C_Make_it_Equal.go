package main

import (
	"bufio"
	"fmt"
	"os"
)

var in *bufio.Reader
var out *bufio.Writer

func solve() {
	var n, k int
	fmt.Fscan(in, &n, &k)

	s := make([]int, n)
	t := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &s[i])
	}

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &t[i])
	}

	m := make(map[int]int)
	for i := 0; i < n; i++ {
		s[i] %= k
		t[i] %= k

		m[s[i]]++
		m[k-s[i]]++
		m[t[i]]--
		m[k-t[i]]--
	}

	for _, v := range m {
		if v != 0 {
			fmt.Fprintln(out, "NO")
			return
		}
	}
	fmt.Fprintln(out, "YES")
}

func main() {
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var tt int
	fmt.Fscan(in, &tt)
	for ; tt > 0; tt-- {
		solve()
	}
}
