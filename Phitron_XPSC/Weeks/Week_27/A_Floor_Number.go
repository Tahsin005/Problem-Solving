package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var t int

func main() {
	var input io.Reader = bufio.NewReader(os.Stdin)
	fmt.Fscan(input, &t)

	for o := 0; o < t; o++ {

		var n, x int

		fmt.Fscan(input, &n)
		fmt.Fscan(input, &x)
		if n == 1 || n == 2 {
			fmt.Println(1)
		} else {
			k := n - 2
			m := k % x
			if m == 0 {
				fmt.Println(int(k / x) + 1)
			} else {
				fmt.Println(int(k / x) + 2)
			}
		}

	}
}