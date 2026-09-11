package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	days := n + (n - 1) / (m - 1)
	fmt.Print(days)
}
