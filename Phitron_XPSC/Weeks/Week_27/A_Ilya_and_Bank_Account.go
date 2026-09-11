package main

import "fmt"

func main () {
	var n int
	fmt.Scan(&n)

	lastRemove := n / 10
	secondLastRemove := n / 100 * 10 + n % 10

	fmt.Println(max(n, lastRemove, secondLastRemove))
}