package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)
	for t > 0 {
		t--
		var n, m, p, q int64
		fmt.Fscan(in, &n, &m, &p, &q)

		groups := n / p
		rem := n % p

		if rem == 0 {
			if m == groups * q {
				fmt.Fprintln(out, "YES")
			} else {
				fmt.Fprintln(out, "NO")
			}
		} else {
			fmt.Fprintln(out, "YES")
		}
	}
}
