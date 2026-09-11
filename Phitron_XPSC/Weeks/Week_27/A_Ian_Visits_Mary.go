package main

import (
	"fmt"
)

func solve() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(2)
	fmt.Println(a - 1, 1)
	fmt.Println(a, b)
}

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		solve()
	}
}
