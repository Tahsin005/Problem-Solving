package main

import (
	."fmt"
	"os"
	"bufio"
)

var in = bufio.NewReader(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solve() {
	var (
		n int
		s string
	)
	Fscan(in, &n, &s)
	ans := make([]int, n)
	val:= 1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '<' {
			ans[i + 1] = val
			val++
		} else {
			ans[i + 1] = n
			n--
		}
	}
	ans[0] = n
	for _, value := range ans {
		Fprint(out, value, " ")
	}
	Fprintln(out, )
}

func main() {

	defer out.Flush()

	var T int
	for Fscan(in, &T); T > 0; T-- {
		solve()
	}
}