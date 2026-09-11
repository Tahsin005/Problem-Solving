package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n, m int
		fmt.Scan(&n, &m)

		sum_a := 0
		for j := 0; j < n; j++ {
			var a int
			fmt.Scan(&a)
			sum_a += a
		}

		ans := sum_a - m
		if ans < 0 {
			ans = 0
		}
		fmt.Println(ans)
	}
}