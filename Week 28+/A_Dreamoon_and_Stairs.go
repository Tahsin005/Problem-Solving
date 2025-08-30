package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	d := (n + 1) / 2
	d = (d + m - 1) / m * m
	if d <= n {
		fmt.Println(d)
	} else {
		fmt.Println(-1)
	}
}
