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
	var a string
	fmt.Fscan(in, &a)
	av := make(map[int]int)
	for i := 0; i < len(a); i++ {
		av[int(a[i] - '0')]++
	}
	for i := 0; i < len(a); i++ {
		for j := 9 - i; j < 10; j++ {
			if av[j] > 0 {
				av[j]--

				fmt.Fprint(out, j)
				break
			}
		}
	}
	fmt.Fprintln(out)
}
func main () {
	defer out.Flush()
	var testCases int
	fmt.Fscan(in, &testCases)

	for testCases > 0 {
		solve()
		testCases--
	}
}