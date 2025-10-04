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
		mp := make(map[int]int)

		for i := 0; i < n; i++ {
			var temp int
			fmt.Fscan(in, &temp)
			mp[temp]++
		}

		ans := len(mp) * 2 - 1

		fmt.Fprintln(out, ans)
	}
}
