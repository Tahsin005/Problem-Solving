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
		var n, m int
		fmt.Fscan(in, &n, &m)

		a := make([]int, m)
		for i := 0; i < m; i++ {
			fmt.Fscan(in, &a[i])
		}

		maxNum := a[0]
		for _, num := range a {
			if num > maxNum {
				maxNum = num
			}
		}

		if maxNum != a[m - 1] || m > maxNum {
			fmt.Fprintln(out, 1)
			continue
		}

		ans := n - maxNum + 1
		fmt.Fprintln(out, ans)
	}
}
