package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	for ; t > 0; t-- {
		var n int
		fmt.Scan(&n)

		if n % 4 == 0 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
