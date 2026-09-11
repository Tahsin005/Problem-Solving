package main

import "fmt"

func main () {
	var testCases int
	fmt.Scan(&testCases)

	for testCases > 0 {
		var n, x int
		fmt.Scan(&n, &x)

		for i := 0; i < n; i++ {
			if i == x {
				continue
			} else {
				fmt.Print(i, " ")
			}
		}
		if x < n {
			fmt.Print(x)
		}
		fmt.Println()
		testCases--
	}
}