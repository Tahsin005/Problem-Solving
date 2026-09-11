package main

import "fmt"

func solve() {
	var t int
	fmt.Scan(&t)
	for ; t > 0; t-- {
		var n, a, b int
		fmt.Scan(&n, &a, &b)
		for i := 0; i < n; i++ {
			fmt.Printf("%c", 'a' + i % b)
		}
		fmt.Println()
	}
}
func main() {
	solve()
}
