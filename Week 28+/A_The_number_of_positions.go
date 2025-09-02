package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)

	fmt.Println(min(n - a, b + 1))
}