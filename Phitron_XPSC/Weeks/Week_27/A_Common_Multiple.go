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
	var t, n, num int
	fmt.Fscan(in, &t)

	for range t {
		fmt.Fscan(in, &n)
		mp := map[int]bool{}

		res := 0

		for range n {
			fmt.Fscan(in, &num)
			_, ok := mp[num]
			if !ok {
				res++
				mp[num] = true
			}
		}

		fmt.Println(res)
	}
}