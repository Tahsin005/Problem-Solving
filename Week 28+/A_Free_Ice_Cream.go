package main

import "fmt"

func main() {
	var n, m, x, i, cnt int64
	fmt.Scan(&n, &m)
	var s string
	for i = 0; i < n; i++ {
		fmt.Scan(&s, &x)
		if s == "+" {
			m += x
		} else if s == "-" && m - x >= 0 {
			m -= x
		} else {
			cnt++
		}
	}
	fmt.Print(m, " ", cnt)
}
