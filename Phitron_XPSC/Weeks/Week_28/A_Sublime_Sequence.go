package main

import "fmt"

func main () {
	var tt int
	fmt.Scan(&tt)

	for ; tt > 0; tt-- {
		var x, n int
		fmt.Scan(&x, &n)

		if n % 2 == 0 {
			fmt.Println(0)
		} else {
			fmt.Println(x)
		}
	}
}
