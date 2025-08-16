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
	for ; t > 0; t-- {
		var n int
		fmt.Fscan(in, &n)
		val, ok := -1, true
		for i := 0; i < n; i++ {
			var x int
			fmt.Fscan(in, &x)
			if x == 0 {
				ok = false
			} else if x != -1 {
				if val == -1 {
					val = x
				} else if val != x {
					ok = false
				}
			}
		}
		if ok {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}
