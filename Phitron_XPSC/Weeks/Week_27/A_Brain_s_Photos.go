package main

import (
	"fmt"
	"strings"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	var color string
	for i := 0; i < n*m; i++ {
		fmt.Scan(&color)
		if strings.ContainsAny(color, "CMY") {
			fmt.Println("#Color")
			return
		}
	}

	fmt.Println("#Black&White")
}
