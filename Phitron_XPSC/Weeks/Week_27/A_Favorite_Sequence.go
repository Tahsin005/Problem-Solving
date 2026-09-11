package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t)
	for t > 0 {
		var n int
		fmt.Scan(&n)
		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&a[i])
		}

		l, r := 0, n
		flag := true
		for l < r {
			if flag {
				fmt.Printf("%d ", a[l])
				l++
			} else {
				fmt.Printf("%d ", a[r-1])
				r--
			}
			flag = !flag
		}
		fmt.Println()

		t--
	}
}
