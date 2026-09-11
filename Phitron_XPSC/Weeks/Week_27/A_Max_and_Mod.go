package main

import (
	"fmt"
)

func main() {

	var t int
	fmt.Scan(&t)
	
	for ; t > 0; t-- {
		var n int
		fmt.Scan(&n)

		if n & 1 == 1 {
			fmt.Print(n)
			for i := 1; i < n; i++ {
				fmt.Print(" ", i)
			}
			fmt.Println()
		} else {
			fmt.Println(-1)
		}
	}
}
