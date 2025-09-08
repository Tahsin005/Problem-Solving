package main

import "fmt"

func main () {
	var tt int
	fmt.Scan(&tt)

	for tt > 0 {
		tt--
		var k, x int
		fmt.Scan(&k, &x)

		for i := 1; i <= k; i++ {
			x *= 2
		}
		fmt.Println(x)
	}
}