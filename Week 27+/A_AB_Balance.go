package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://codeforces.com/problemset/problem/1606/A

var (
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func main() {
	defer out.Flush()

	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var s string
		fmt.Scan(&s)

		if len(s) > 0 {
			lastChar := string(s[len(s) - 1])
			rest := s[1:]
			fmt.Println(lastChar + rest)
		}
	}
}
