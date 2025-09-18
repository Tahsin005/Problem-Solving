package main

import "fmt"

func main() {
	var tt int
	fmt.Scan(&tt)

	for t := 1; t <= tt; t++ {
		var n int
		fmt.Scan(&n)

		for i := 2; i <= n; i++ {
			fmt.Print(i, " ")
		}
		fmt.Println(1)
	}
}
