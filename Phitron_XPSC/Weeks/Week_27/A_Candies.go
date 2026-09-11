package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n)

		k := 3
		for k <= n {
			if n % k == 0 {
				fmt.Println(n / k)
				break
			}

			k = 2 * k + 1
		}
	}
}