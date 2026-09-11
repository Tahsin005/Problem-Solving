package main

import "fmt"

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func main () {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		minS := min(a, b)
		maxS := max(a, b)
		s := max(2 * minS, maxS)
		fmt.Println(s * s)
	}
}