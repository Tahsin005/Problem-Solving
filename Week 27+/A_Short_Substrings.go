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

func main() {
	defer out.Flush()

	var testCases int
	fmt.Fscan(in, &testCases)

	for i := 0; i < testCases; i++ {
		var s string
		fmt.Fscan(in, &s)

		ans := ""
		ans += string(s[0])
		for j := 1; j < len(s); j += 2 {
			ans += string(s[j])
		}
		fmt.Fprintln(out, ans)
	}
}
