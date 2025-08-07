package main

import "fmt"

func main() {
	t := 1
	fmt.Scan(&t)
	for k := 0; k < t; k++ {
		x0, sum := 0, 0
		var n, x int
		fmt.Scan(&n)
		for i := 0; i < n; i++ {
			fmt.Scan(&x)
			if x == 0 {
				x0++
			}
			sum += x
		}
		fmt.Println(sum + x0)
	}
}
