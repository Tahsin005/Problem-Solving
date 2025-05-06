package main

import (
	"fmt"
)

func main() {
	var a [4]int
	for i := range a {
		fmt.Scan(&a[i])
	}

	var s string
	fmt.Scan(&s)
	z := 0
	for _, c := range s {
		z += a[int(c - '1')]
	}
	fmt.Println(z)
}
