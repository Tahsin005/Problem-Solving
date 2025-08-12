package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n)

		series := make([]int, n)

		for i := 0; i < n; i++ {
			if i % 2 == 0 {
				series[i] = -1
			} else {
				series[i] = 3
			}
		}

		if n % 2 == 0 {
			series[n - 1] = 2
		}
		for _, v := range series {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}
}