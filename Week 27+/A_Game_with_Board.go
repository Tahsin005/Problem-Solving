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
	defer out.Flush()

	var testCases int
	fmt.Scan(&testCases)

	for i := 0; i < testCases; i++ {
		var n int
		fmt.Scan(&n)

		if n >= 5 {
			fmt.Println("Alice")
		} else {
			fmt.Println("Bob")
		}
	}
}