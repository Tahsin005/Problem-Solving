package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n)
		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&a[i])
		}
		fmt.Println(a[0])
	}
}
