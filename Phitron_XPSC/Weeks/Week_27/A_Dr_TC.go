package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func solve() {
	var n int
	fmt.Fscan(in, &n)
	var s string
	fmt.Fscan(in, &s)

	ans := 0
	for i := range n {
		if s[i] == '0' {
			ans++
		} else {
			ans += (n - 1)
		}
	}

	fmt.Fprintln(out, ans)
}

func main () {
	defer out.Flush()
	var testCases int
	fmt.Fscan(in, &testCases)
	for i := 0; i < testCases; i++ {
		solve()
	}
}