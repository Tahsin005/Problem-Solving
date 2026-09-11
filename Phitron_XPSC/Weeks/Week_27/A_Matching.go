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
	fmt.Fscan(in, &testCases)
	for i := 0; i < testCases; i++ {
		var s string
		fmt.Fscan(in, &s)

		ans := 1

		if s[0] == '0' {
			ans *= 0
		} else {
			if s[0] == '?' {
				ans *= 9
			}

			for n := 1; n < len(s); n++ {
				if s[n] == '?' {
					ans *= 10
				}
			}
		}

		fmt.Println(ans)
	}
}