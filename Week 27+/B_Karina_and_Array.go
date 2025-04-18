package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func solve () {
	var n int
	fmt.Fscan(in, &n)

	a := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	sort.Ints(a)

	ans := int64(a[0] * a[1])

	if ans < int64(a[n - 1] * a[n - 2]) {
		ans = int64(a[n - 1] * a[n - 2])
	}
	fmt.Println(ans)
}

func main() {
	defer out.Flush()
	var testCases int
	fmt.Fscan(in, &testCases)

	for i := 0; i < testCases; i++ {
		solve()
	}
}