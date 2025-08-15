package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer

	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		var n, k int
		fmt.Fscan(in, &n)
		fmt.Fscan(in, &k)
		towers := make([]int, 0)
		for ; n > 0; n-- {
			var height int
			fmt.Fscan(in, &height)
			towers = append(towers, height)
		}
		curr := towers[k - 1]
		height := 1
		sort.Slice(towers, func(i, j int) bool {
			return towers[i] < towers[j]
		})
		ans := true
		for i := 0; i < len(towers); i++ {
			if towers[i] < curr {
				continue
			}
			if towers[i] - curr > curr + 1 - height {
				ans = false
				break
			}
			height += towers[i] - curr
			curr = towers[i]
		}
		if ans {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}
