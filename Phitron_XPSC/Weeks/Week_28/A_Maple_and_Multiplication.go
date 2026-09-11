package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		if a == b {
			fmt.Println(0)
		} else if (a % b == 0 || b % a == 0) {
			fmt.Println(1)
		} else {
			fmt.Println(2)
		}
	}
}
