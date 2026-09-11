package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var i int
	for r, s := 0, 0; s <= n; s += r {
		i++
		r += i
	}
	i--
	fmt.Println(i)
}
