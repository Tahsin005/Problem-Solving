package main

import"fmt"

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n)
		for j := 0; j < n; j++ {
			var num int
			fmt.Scan(&num)
			fmt.Printf("%d ", n + 1 - num)
		}
		fmt.Println()
	}
}
