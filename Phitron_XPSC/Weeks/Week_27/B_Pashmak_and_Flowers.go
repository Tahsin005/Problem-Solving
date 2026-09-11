package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func solve() {
	var n int

	fmt.Fscan(in, &n)

	arr := make([]int, n)

	mn := math.MaxInt
	mx := math.MinInt

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
		if arr[i] > mx {
			mx = arr[i]
		}
		if arr[i] < mn {
			mn = arr[i]
		}
	}

	c1 := 0
	c2 := 0

	for i := 0; i < n; i++ {
		if arr[i] == mn {
			c1++
		}

		if arr[i] == mx {
			c2++
		}
	}

	if mn == mx {
		fmt.Fprintln(out, mx - mn, n * (n - 1) / 2)
	} else {
		fmt.Fprintln(out, mx - mn, c1 * c2)
	}
}

func main() {
	defer out.Flush()
	solve()
}
