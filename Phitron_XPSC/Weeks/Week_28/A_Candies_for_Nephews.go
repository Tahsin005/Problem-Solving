package main

import "fmt"

func main () {
	var tt int
	fmt.Scan(&tt)
	for t := 0; t < tt; t++ {
		var n int
		fmt.Scan(&n)

		if (n % 3 == 0) {
			fmt.Println(0)
		} else {
			fmt.Println(3 - (n % 3))
		}
	}
}