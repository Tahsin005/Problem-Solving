package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		var n, k int
		fmt.Fscan(in, &n, &k)
		if k * k > n || k * k % 2 != n % 2 {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
	}
}
