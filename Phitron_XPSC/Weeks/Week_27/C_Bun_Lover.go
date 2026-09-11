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

func main() {
	defer out.Flush()

	var testCases int
	var n int

	fmt.Fscan(in, &testCases)

	for i := 0; i < testCases; i++ {
		fmt.Fscan(in, &n)

		n++
		totalLengthOfChocolate := n * n + 1

		fmt.Fprintln(out, totalLengthOfChocolate)
	}
}