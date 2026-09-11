package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func solve() {
	var n int
	fmt.Fscan(in, &n)

	if n % 2 == 0 {
		fmt.Println("-1")
		return
	}

	var res []int
	for n > 1 {
		if ((n - 1) / 2) % 2 != 0 {
			res = append(res, 2)
			n = (n - 1) / 2
		} else {
			res = append(res, 1)
			n = (n + 1) / 2
		}
	}

	fmt.Println(len(res))

	for j := len(res) - 1 ; j >= 0 ; j -- {
		fmt.Printf("%d ", res[j])
	}
	
	fmt.Println("")
}

func main() {
	defer out.Flush()

	var testCases int
	fmt.Fscan(in, &testCases)

	for testCases > 0 {
		solve()
		testCases--
	}
}
