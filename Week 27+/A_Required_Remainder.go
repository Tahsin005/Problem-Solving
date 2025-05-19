package main

import (
	"bufio"
	"fmt"
	"os"
)

var read *bufio.Reader = bufio.NewReader(os.Stdin)
var write *bufio.Writer = bufio.NewWriter(os.Stdout)

func main()  {
	defer write.Flush()
	var t int
	fmt.Fscan(read, &t)
	for ; t > 0; t-- {
		var x, y, n int
		fmt.Fscan(read, &x, &y, &n)
		k := (n - y) / x
		k = k * x + y
		fmt.Fprintln(write, k)
	}
}