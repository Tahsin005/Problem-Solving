package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func main () {
	var testCases int
	fmt.Fscan(in, &testCases)
	for i := 0; i < testCases; i++ {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)

		total := a + b + c

		if total % 3 != 0 || a + c < 2 * b {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
	}
}