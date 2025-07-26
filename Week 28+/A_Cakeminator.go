package main

import (
	"fmt"
)

func main() {
	var r, c int
	cs := 0
	var s string
	fmt.Scan(&r, &c)
	rw := make([]int, r)
	cl := make([]int, c)
	for i := 0; i < r; i++ {
		fmt.Scan(&s)
		for j := 0; j < c; j++ {
			if s[j] == 'S' {
				rw[i] = 1
				cl[j] = 1
			}
		}
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if rw[i] == 0 || cl[j] == 0 {
				cs++
			}
		}
	}
	fmt.Println(cs)
}
