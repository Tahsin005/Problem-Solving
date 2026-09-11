package main

import "fmt"

func main() {
	var n, m, a int
	fmt.Scan(&n, &m)
	var lastOne, t int
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		x := (a + m - 1) / m
		if x >= t {
			t = x
			lastOne = i
		}
	}
	fmt.Println(lastOne + 1)
}
